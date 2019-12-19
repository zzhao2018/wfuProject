package etcdRegister

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"path"
	"sync"
	"sync/atomic"
	"time"
	"wfuProject/register"
)

//etcd注册引擎
type EtcdRegister struct {
	cli               *clientv3.Client               //etcd客户端
	serverChan        chan *register.Server          //保存未处理的server
	serverRegisterMap map[string]*ServerRegisterMeta //保存服务器结点信息
	options           *register.RegisterOption       //保存etcd配置
	serverCacheMap    atomic.Value                   //用于保存缓存的数据
	lock              sync.Mutex
}

//用以保存每个server元数据
type ServerRegisterMeta struct {
	server         *register.Server
	registerCh     <-chan *clientv3.LeaseKeepAliveResponse
	registerStatus bool
}

//保存缓存的server数据
type cacheServerData struct {
	serverMap map[string]*register.Server
}

var (
	E_PullCacheTime = time.Second * 5
	etcdRegister    = &EtcdRegister{
		serverChan:        make(chan *register.Server, 8),
		serverRegisterMap: make(map[string]*ServerRegisterMeta, 8),
		lock:              sync.Mutex{},
	}
)

//加载模块时，初始化引擎
func init() {
	//将模块注册到register
	err := register.RegisterServer(etcdRegister)
	if err != nil {
		log.Printf("etcd_register init register error,err:%+v\n", err)
		return
	}
	//初始化缓存
	cacheServer := cacheServerData{serverMap: make(map[string]*register.Server, 8)}
	etcdRegister.serverCacheMap.Store(cacheServer)
	//开启后台协程
	go etcdRegister.run()
}

/********************定时pull服务信息至cache*********************/
func (e *EtcdRegister) pullToUpdateCache() {
	oldMap := e.serverCacheMap.Load().(cacheServerData)
	newMap := cacheServerData{serverMap: make(map[string]*register.Server)}
	for serverName, serverData := range oldMap.serverMap {
		//从etcd获得数据
		name := path.Join(e.options.RegisterPath, serverName)
		respData, err := e.cli.Get(context.TODO(), name, clientv3.WithPrefix())
		if err != nil {
			log.Printf("pullToUpdateCache get data error,err:%+v\n", err)
			newMap.serverMap[serverName] = serverData
			continue
		}
		//更新
		changeError := false
		tempServer := &register.Server{
			Name: serverName,
			Node: make([]*register.ServerNode, 0),
		}
		for _, respEle := range respData.Kvs {
			tempServerEle := &register.Server{}
			err = json.Unmarshal(respEle.Value, tempServerEle)
			if err != nil {
				log.Printf("pullToUpdateCache unmarshal error,err:%+v\n", err)
				newMap.serverMap[serverName] = serverData
				changeError = true
				break
			}
			tempServer.Node = append(tempServer.Node, tempServerEle.Node...)
		}
		if changeError == false {
			newMap.serverMap[serverName] = tempServer
		}
	}
	//更新
	e.serverCacheMap.Store(newMap)
}

/********************注册操作********************/
//注册节点至etcd
func (e *EtcdRegister) doRegister(serverData *ServerRegisterMeta) error {
	//初始化续租id
	grantid, err := e.cli.Grant(context.TODO(), e.options.HeartBeat)
	if err != nil {
		log.Printf("doRegister grand id error,err:%+v\n", err)
		return err
	}
	//注册
	//key:/registerpath/server名称/nodeip:nodeport
	//value: server的json
	for _, serverNode := range serverData.server.Node {
		temp := &register.Server{
			Name: serverData.server.Name,
			Node: []*register.ServerNode{serverNode},
		}
		key := e.getKey(temp)
		log.Printf("register key:%s\n",key)
		value, err := json.Marshal(temp)
		if err != nil {
			log.Printf("doRegister marshal error,err:%+v\n", err)
			continue
		}
		_, err = e.cli.Put(context.TODO(), key, string(value), clientv3.WithLease(grantid.ID))
		if err != nil {
			log.Printf("doRegister put data error,(key:%+v value:%+v),err:%+v\n", key, value, err)
			continue
		}
		//续期
		ch, err := e.cli.KeepAlive(context.TODO(), grantid.ID)
		if err != nil {
			log.Printf("doRegister keepalive error,(key:%+v value:%+v),err:%+v\n", key, value, err)
			continue
		}
		serverData.registerCh = ch
		serverData.registerStatus = true
	}
	return nil
}

//获得server的key
func (e *EtcdRegister) getKey(server *register.Server) string {
	nodePath := fmt.Sprintf("%s:%s", server.Node[0].Ip, server.Node[0].Port)
	return path.Join(e.options.RegisterPath, server.Name, nodePath)
}

//检查是否有未注册的服务，有则开启服务
func (e *EtcdRegister) CheckServer() {
	for _, value := range e.serverRegisterMap {
		//已经注册的保持活性
		if value.registerStatus == true {
			e.keepAlive(value)
			continue
		}
		//失活的注册
		e.doRegister(value)
	}
}

//用以定时发送数据
func (e *EtcdRegister) keepAlive(meta *ServerRegisterMeta) {
	select {
	case value := <-meta.registerCh:
		//若服务注册失效，则重新注册
		if value == nil {
			meta.registerStatus = false
		}
		//log.Printf("chan get value:%+v\n",value)
	}
}

//后台协程，用以注册服务
//若有服务需要注册，则放入map中
//否则检查是否有需要注册的服务
func (e *EtcdRegister) run() {
	timeTick := time.NewTicker(E_PullCacheTime)
	for {
		//从chan中取出未注册的服务，注册
		select {
		case server := <-e.serverChan:
			//判断server是否存在
			if oldServer, ok := e.serverRegisterMap[server.Name]; ok == true {
				//服务名称相同时，增加结点，不覆盖
				for _, ele := range server.Node {
					oldServer.server.Node = append(oldServer.server.Node, ele)
				}
				e.serverRegisterMap[server.Name].registerStatus = false
				//log.Printf("do register fail,%s exist\n", server.Name)
				break
			}
			//初始化
			serverData := &ServerRegisterMeta{
				server:         server,
				registerStatus: false,
			}
			e.serverRegisterMap[server.Name] = serverData
		case <-timeTick.C:
			//定时更新缓存
			e.pullToUpdateCache()
		default:
			//检查是否有未注册的服务，有则开启服务
			e.CheckServer()
		}
	}
}

/********************是实现register接口的方法*****************/
func (e *EtcdRegister) Name() string {
	return "etcd"
}

/////etcd初始化
func (e *EtcdRegister) Init(ctx context.Context, opts ...register.RegisterOptionFunc) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	//初始化配置
	opt := &register.RegisterOption{}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	e.options = opt
	//设置etcd
	var err error
	e.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   opt.Addr,
		DialTimeout: opt.TimeOut,
	})
	if err != nil {
		log.Printf("Init error,init etcd error,err:%+v\n", err)
		return err
	}
	return nil
}

/////服务注册
//将服务放置入chan后即可
func (e *EtcdRegister) Register(ctx context.Context, server *register.Server) error {
	select {
	case e.serverChan <- server:
	default:
		log.Printf("Register error,server channel is full\n")
		err := fmt.Errorf("server channel is full")
		return err
	}
	return nil
}

//服务反注册
func (e *EtcdRegister) UnRegister(ctx context.Context, server *register.Server) error {
	return nil
}

/////服务拉取
//从缓存中查询数据
func (e *EtcdRegister) getServerFromCache(serverName string) (*register.Server, bool) {
	cacheMap := e.serverCacheMap.Load().(cacheServerData)
	value, ok := cacheMap.serverMap[serverName]
	return value, ok
}

//拉取服务列表
func (e *EtcdRegister) GetServer(ctx context.Context, serverName string) (*register.Server, error) {
	//查询缓存
	serverData, ok := e.getServerFromCache(serverName)
	if ok == true {
		//log.Printf("hit cache,data:%+v\n", serverData)
		return serverData, nil
	}
	//若未命中缓存，查询etcd
	e.lock.Lock()
	defer e.lock.Unlock()
	//再次查询cache
	serverData, ok = e.getServerFromCache(serverName)
	if ok == true {
		//log.Printf("hit cache,data:%+v\n", serverData)
		return serverData, nil
	}
	//构造key
	key := path.Join(e.options.RegisterPath, serverName)
	respData, err := e.cli.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		log.Printf("GetServer get data error,err:%+v\n", err)
		return nil, err
	}
	//从etcd中查询数据
	tempServer := &register.Server{
		Name: serverName,
		Node: make([]*register.ServerNode, 0),
	}
	for _, ele := range respData.Kvs {
		tempServerEle := &register.Server{}
		err = json.Unmarshal(ele.Value, tempServerEle)
		if err != nil {
			log.Printf("GetServer unmarsh error,err:%+v\n", err)
			return nil, err
		}
		tempServer.Node = append(tempServer.Node, tempServerEle.Node...)
	}
	//更新缓存
	changeFlag := false
	if len(respData.Kvs) > 0 {
		changeFlag = true
	}
	if changeFlag {
		oldServerData := e.serverCacheMap.Load().(cacheServerData)
		oldServerData.serverMap[serverName] = tempServer
		e.serverCacheMap.Store(oldServerData)
	}
	return tempServer, nil
}

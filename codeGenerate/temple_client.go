package main


var templeClientMain=`
package main

import (
)

func main() {
}
`

var templeClientService=`
package clientService

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"wfuProject/clientMidware"
	"wfuProject/codeGenerate/{{.OutputPath}}/generate"
)

const(
	addr ="localhost:12306"
)

type OptionFunc func(c *ClientDeal)

type ClientDeal struct {
	serverName string
	traceid string
}

/**************初始化*****************/
func NewClientDeal(serverName string,opts ...OptionFunc)*ClientDeal{
	client:=&ClientDeal{
		serverName:serverName,
	}
	for _,eleFunc:=range opts {
		eleFunc(client)
	}
	return client
}

func OptClientTraceId(traceids string)OptionFunc{
	return func(c *ClientDeal) {
		c.traceid=traceids
	}
}

/*******************方法封装******************/
{{range .Rpc}}
func(c *ClientDeal){{.Name}}(ctx context.Context, in *generate.{{.RequestType}}, opts ...grpc.CallOption) (*generate.{{.ReturnsType}}, error) {
	//设置分布式追踪traceid
	ctx=c.contextWithTraceMidWareValue(ctx)
	//启动中间件
	outFunc:=clientMidware.BuildClientMidWareLink({{.Name}}ClientHandler)
	response,err:=outFunc(ctx,in)
	if err!=nil {
		log.Printf("{{.Name}} getResponse error,err:%+v\n",err)
		return nil,err
	}
	responseData,ok:=response.(*generate.{{.ReturnsType}})
	if ok==false {
		err=fmt.Errorf("{{.Name}} change response error,err:%+v\n",err)
		log.Printf("{{.Name}} change response error\n")
		return nil,err
	}
	return responseData,nil
}

func {{.Name}}ClientHandler(ctx context.Context,request interface{})(interface{},error){
	//连接grpc
	conn,err:=grpc.Dial(addr,grpc.WithInsecure())
	if err!=nil {
		log.Printf("{{.Name}}ClientHandler dial grpc error,err:%+v\n",err)
		return nil,err
	}
	defer conn.Close()
	client:=generate.New{{$.Service.Name}}Client(conn)
	//获得request
	requestClient,ok:=request.(*generate.{{.RequestType}})
	if ok==false {
		err=fmt.Errorf("change request type error")
		log.Printf("{{.Name}}ClientHandler change type error,err:%+v\n",err)
		return nil,err
	}
	response,err:=client.{{.Name}}(ctx,requestClient)
	if err!=nil {
		log.Printf("{{.Name}}ClientHandler {{.Name}} error,err:%+v\n",err)
		return nil,err
	}
	return response,nil
}
{{end}}

/********context携带traceid*******/
func(c *ClientDeal) contextWithTraceMidWareValue(ctx context.Context)context.Context{
	ctx=context.WithValue(ctx,clientMidware.ClientTraceServerName{},c.serverName)
	return context.WithValue(ctx,clientMidware.ClientTraceIdKey{},c.traceid)
}

`
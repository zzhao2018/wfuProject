package main


var templeClientMain=`
package main

import (
)

func main() {
}
`

var templeClientService=`
package clientTool

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"wfuProject/clientService"
    "wfuProject/clientUtil"
	"wfuProject/codeGenerate/{{.OutputPath}}/generate"
)

const(
	addr ="localhost:12306"
)

type ClientDeal struct {
	clientService *clientService.ClientService
}


/**************初始化*****************/
func NewClientDeal(serverName string,opts ...clientService.OptionFunc)*ClientDeal{
	cs:=clientService.NewClientService(serverName,opts...)
	return &ClientDeal{clientService:cs}
}

/*******************方法封装******************/
{{range .Rpc}}
func(c *ClientDeal){{.Name}}(ctx context.Context, in *generate.{{.RequestType}}, opts ...grpc.CallOption) (*generate.{{.ReturnsType}}, error) {
	//调用call
    response,err:=c.clientService.Call(ctx,{{.Name}}ClientHandler,in,"{{.Name}}")
	if err!=nil {
		log.Printf("{{.Name}} Call error,err:%+v\n",err)
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
	//从短连接中获取conn
	serverMetaData,err:=clientUtil.GetMetaDataFromContext(ctx)
	if err!=nil {
		log.Printf("client_tool SumClientHandler GetMetaDataFromContext error,err:%+v\n",err)
		return nil,err
	}
	conn:=serverMetaData.Conn
	client:=generate.NewTestClient(conn)
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
`
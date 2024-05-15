package main

import (
	"fmt"
	apiServices "jrpc/src/rpc_api/services"
	transportClient "jrpc/src/rpc_core/transport/client"
)

func main() {
	// client := transportClient.NewDefaultSocketClient()
	client := transportClient.NewDefaultSocketClientWithAimIp("127.0.0.1:3212")
	proxy := transportClient.NewRpcClientProxy(client)

	testService := proxy.NewProxyInstance(&apiServices.Test{}).(*apiServices.Test)
	res, _ := testService.Ping()
	fmt.Println("结果是: ", res)
	res, _ = testService.Hello()
	fmt.Println("结果是: ", res)
}
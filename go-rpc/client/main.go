package main

import (
	"fmt"
	"go_rpc/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string //string有默认值

	msg := "Hello"
	err := client.Echo(msg, &reply)
	fmt.Println("client=> ", msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("server=> ", reply)
}

package main

import (
	hanlder "go_rpc/handler"
	"go_rpc/server_proxy"
	"net"
	"net/rpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	_ = server_proxy.RegisterHelloService(&hanlder.NewEchoService{})
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}

}

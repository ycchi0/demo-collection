package server_proxy

import (
	hanlder "go_rpc/handler"
	"net/rpc"
)

// service interface
type EchoServicer interface {
	Echo(request string, reply *string) error
}

func RegisterHelloService(srv EchoServicer) error {
	return rpc.RegisterName(hanlder.EchoServiceName, srv)
}

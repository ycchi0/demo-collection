package client_proxy

import (
	hanlder "go_rpc/handler"
	"net/rpc"
)

type EchoServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protcol, address string) EchoServiceStub {
	conn, err := rpc.Dial(protcol, address)
	if err != nil {
		panic("connect error!")
	}
	return EchoServiceStub{conn}
}

func (c *EchoServiceStub) Echo(request string, reply *string) error {

	// server registe name and function
	err := c.Call(hanlder.EchoServiceName+".Echo", request, reply)
	if err != nil {
		return err
	}
	return nil
}

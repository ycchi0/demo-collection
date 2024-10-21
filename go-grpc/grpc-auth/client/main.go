package main

import (
	"context"
	"fmt"
	"go-grpc/grpc-auth/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "hello",
		"appkey": "golang1",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(customCredential{}))
	conn, err := grpc.NewClient("127.0.0.1:50051", opts...)
	if err != nil {
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "golang"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

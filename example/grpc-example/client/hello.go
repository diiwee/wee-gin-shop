package main

import (
	"context"
	"fmt"
	"wee-example/grpc-example/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial(":3222", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	c := proto.NewHelloClient(conn)

	ctx := context.Background()

	resp, _ := c.SayHello(ctx, &proto.HelloRequest{Name: "Hello wee grpc example"})

	fmt.Println(resp)
}

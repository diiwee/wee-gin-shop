package main

import (
	"context"
	"fmt"
	"net"
	"wee-example/grpc-example/proto"

	"google.golang.org/grpc"
)

type helloServer struct {
	proto.UnimplementedHelloServer
}

func (h *helloServer) SayHello(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Printf("This demo from name:%s", req.Name)
	return &proto.HelloResponse{Message: req.Name}, nil
}

func main() {

	lis, _ := net.Listen("tcp", ":3222")

	s := grpc.NewServer()

	proto.RegisterHelloServer(s, &helloServer{})

	s.Serve(lis)
}

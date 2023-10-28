package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"massengerGrpc/internal/controllers"
	"net"
)

type gRPCServer struct {
	gRPC *grpc.Server
}

func New() *gRPCServer {
	return &gRPCServer{
		gRPC: grpc.NewServer(),
	}
}

func (server *gRPCServer) Start() {

	if err := server.configureServices(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("asd")
	l, err := net.Listen("tcp", ":8342")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("asd")
	if err := server.gRPC.Serve(l); err != nil {
		fmt.Println("asd")
		log.Fatal(err)
	}
	fmt.Println("asd")

}

func (server *gRPCServer) configureServices() error {

	controllers.RegisterTesterServer(server.gRPC, &controllers.HelloService{})

	return nil
}

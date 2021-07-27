package main

import (
	"grpctt/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	lis, _ := net.Listen("tcp", ":8081")

	s := user.Server{}

	grpcServer := grpc.NewServer()

	user.RegisterHelloWorldServer(grpcServer, &s)

	err := grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}

}
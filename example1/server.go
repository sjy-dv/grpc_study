package main

import (
	"grpctt/example1/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatal("failed", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
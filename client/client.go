package main

import (
	"context"
	"grpctt/chat"
	"log"

	"google.golang.org/grpc"
)


func main() {
	
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8081", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "hello i am client",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("From server : %s", response.Body)
}
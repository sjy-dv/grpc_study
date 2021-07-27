package main

import (
	"context"
	"grpctt/user"
	"log"

	"google.golang.org/grpc"
)


func main() {

	var conn *grpc.ClientConn

	conn, _ = grpc.Dial(":8081", grpc.WithInsecure())

	defer conn.Close()

	c := user.NewHelloWorldClient(conn)

	message := user.Message{
		Body : "test message",
	}

	useridx := user.UserIdx{
		Idx: 1,
	}

	resp, _ := c.SayHello(context.Background(), &message)
	ddew, _ := c.GetUser(context.Background(),&useridx)
	log.Printf(resp.Body)
	log.Print(ddew)
}
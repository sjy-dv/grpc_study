package main

import (
	"grpctt/user"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

//var server = user.DBServer{}


func main() {
	var err error
	
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("env error : %v", err)
	}
//DB_DRIVER,MYSQL_DB_USER, MYSQL_DB_PASSWORD,DB_PORT, MYSQL_DB_HOST,MYSQL_DB
	
	server := user.Server{}

	server.InitDB(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	lis, _ := net.Listen("tcp", ":8081")

	
	grpcServer := grpc.NewServer()

	user.RegisterRpcAppServer(grpcServer, &server)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}
}
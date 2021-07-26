- setpath:
	@export PATH="$PATH:$(go env GOPATH)/bin"

- build:
	protoc --go_out=plugins=grpc:chat chat.proto

- run:
	go run server.go
	go run client/client.go
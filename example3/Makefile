- setpath:
	@export PATH="$PATH:$(go env GOPATH)/bin"

- build:
	protoc --go_out=plugins=grpc:user user.proto

- run:
	go run main.go
	go run client/client.go
create: deps
	protoc --proto_path=proto proto/*.proto --go_out=gen/netvuln
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/netvuln

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: build server client
build: server client
server:
	go build -o ./bin/server ./cmd/server

client:
	go build -o ./bin/client ./cmd/client


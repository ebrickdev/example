
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
.PHONY: proto
proto:
	protoc --proto_path=user --go_out=:user --go-grpc_out=user user/proto/user.proto

.PHONY: build
build:
	go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover
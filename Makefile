SHELL := /bin/bash

init:
	brew install protobuf
	brew install protoc-gen-go
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get google.golang.org/protobuf
	go get google.golang.org/protobuf/proto

migrate:
	protoc --go_out=. *.proto

.PHONY: test
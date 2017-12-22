.PHONY: all

all:

db:
	cd models && go generate

pb:
	protoc -I protos -I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway protos/accounts.proto --go_out=plugins=grpc:pbs

gw:
	protoc -I protos -I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway protos/accounts.proto --grpc-gateway_out=logtostderr=true:pbs
	protoc -I protos -I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway protos/accounts.proto --swagger_out=logtostderr=true:swagger-json

grpc:
	go run main.go grpc

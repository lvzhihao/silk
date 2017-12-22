.PHONY: all

all: gen 
	composer install

gen:
	protoc --proto_path=./../../protos \
	-I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I $$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
	--php_out=./ --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin accounts.proto

run: all
	php client_test.php

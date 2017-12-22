.PHONY: all

all:

gen:
	protoc --proto_path=./../../protos   --php_out=./   --grpc_out=./   --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin  ./../../protos/accounts.proto

run:
	php client_test.php

rpc:
	bash -c "protoc --go_out=. --go-grpc_out=. api/protobuf/*.proto"
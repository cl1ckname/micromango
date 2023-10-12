rpc:
	bash -c "protoc --go_out=. --go-grpc_out=. api/protobuf/*.proto"
	bash -c "protoc-go-inject-tag -input='pkg/grpc/*/*.pb.go'"
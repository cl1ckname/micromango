rpc:
	bash -c "protoc --go_out=. --go-grpc_out=. -I api/protobuf/ --go_opt=module=micromango --go-grpc_opt=module=micromango --go_opt=paths=import api/protobuf/*.proto"
	bash -c "protoc-go-inject-tag -input='pkg/grpc/*/*.pb.go'"
client:
	rm -rf micromango-client; \
	git clone -b main git@github.com:cl1ckname/micromango-client.git;
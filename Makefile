all: rpc client up restart down sql
.PHONY : all

rpc:
	bash -c "protoc --go_out=. --go-grpc_out=. -I api/protobuf/ --go_opt=module=micromango --go-grpc_opt=module=micromango --go_opt=paths=import api/protobuf/*.proto"
client:
	rm -rf micromango-client; \
	git clone -b main git@github.com:cl1ckname/micromango-client.git;

up:
	pm2 start ecosystem.config.js
restart:
	pm2 restart ecosystem.config.js
down:
	pm2 stop ecosystem.config.js
sql:
	sqlite3 db/catalog.sqlite3 < sql/genres.sql
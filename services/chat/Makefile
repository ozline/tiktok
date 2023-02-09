.PHONY: generate
generate:
	kitex \
	-service "tiktok-chat" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ chat.proto \
	&& go mod tidy

.PHONY: run_server
run_mysql_server:
	go run ./mysql_layer/main.go ./mysql_layer/handler.go ./mysql_layer/mysql_handler.go

run_access_server:
	go run ./access_layer/main.go ./access_layer/server.go ./access_layer/user.go ./access_layer/message.go

run_sendClient:
	go run ./client/send.go ./client/message.go

run_receiveClient:
	go run ./client/receive.go ./client/message.go

# Service: chat
# Path: services/chat
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
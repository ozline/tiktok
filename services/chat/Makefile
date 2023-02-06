.PHONY: generate
generate:
	kitex \
	-service "tiktok-chat" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ chat.proto \
	&& go mod tidy

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: chat
# Path: services/chat
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
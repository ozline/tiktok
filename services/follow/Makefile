.PHONY: generate
generate:
	kitex \
	-service "tiktok-follow" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ follow.proto \
	&& go mod tidy

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: follow
# Path: services/follow
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
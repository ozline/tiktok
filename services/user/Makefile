.PHONY: generate
generate:
	kitex \
	-service "tiktok-user" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ user.proto \
	&& go mod tidy

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: user
# Path: services/user
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
.PHONY: generate
generate:
	kitex \
	-service "tiktok-auth" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ auths.proto

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: auth
# Path: services/auth
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
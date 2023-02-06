.PHONY: generate
generate:
	kitex \
	-service "tiktok-comment" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ comment.proto


.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: comment
# Path: services/comment
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
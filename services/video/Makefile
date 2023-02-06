.PHONY: generate
generate:
	kitex \
	-service "tiktok-video" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ video.proto \
	&& go mod tidy

.PHONY: run_client
run_client:
	go run ./client/main.go

.PHONY: run_server
run_server:
	go run main.go handler.go

# Service: video
# Path: services/video
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
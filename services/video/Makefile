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

test_upMultiVideo:
	go run ./client/up_multi_video.go

test_downMultiVideo:
	go run ./client/down_multi_video.go

test_downOneVideo:
	go run ./client/down_one_video.go

test_getOneVideoInfo:
	go run ./client/get_one_video_info.go

test_deleteOneVideo:
	go run ./client/delete_one_video.go
# Service: video
# Path: services/video
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
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
	go run main.go handler.go mysqlHandler.go snowFlake.go storageHandler.go

test_upMultiVideo:
	go run ./client/upMultiVideoTest.go

test_downMultiVideo:
	go run ./client/downMultiVideoTest.go

test_downOneVideo:
	go run ./client/downOneVideoTest.go

test_getOneVideoInfo:
	go run ./client/getOneVideoInfoTest.go

test_deleteOneVideo:
	go run ./client/deleteOneVideoTest.go
# Service: video
# Path: services/video
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
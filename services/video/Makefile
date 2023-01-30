generate:
	kitex \
	-service "tiktok-video" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ video.proto \
	&& go mod tidy

# Service: video
# Path: services/video
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
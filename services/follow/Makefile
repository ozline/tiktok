generate:
	kitex \
	-service "tiktok-follow" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ follow.proto \
	&& go mod tidy

# Service: follow
# Path: services/follow
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
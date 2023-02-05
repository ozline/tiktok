.PHONY: generate
generate:
	kitex \
	-service "tiktok-user" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ user.proto \
	&& go mod tidy

# Service: user
# Path: services/user
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
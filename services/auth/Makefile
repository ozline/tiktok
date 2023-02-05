.PHONY: generate
generate:
	kitex \
	-service "tiktok-auth" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ auth.proto

# Service: auth
# Path: services/auth
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
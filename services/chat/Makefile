generate:
	kitex \
	-service "tiktok-chat" \
	-module "github.com/ozline/tiktok" \
	-type protobuf \
	-I ./../../idl/ chat.proto \
	&& go mod tidy
# Service: chat
# Path: services/chat
# Optional: -use "github.com/ozline/tiktok/kitex_gen"
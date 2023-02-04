init:
	hz new \
	-module github.com/ozline/tiktok/api-gateway \
	-service "api-gateway" \
	-proto_path ./../ \
	-idl ./../idl/gateway/gateway.proto \
	&& go mod tidy

update:
	hz update \
	-proto_path ./../ \
	-idl ./../idl/gateway/gateway.proto
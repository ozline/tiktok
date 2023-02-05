.PHONY: init
init:
	hz new \
	-module github.com/ozline/tiktok/api-gateway \
	-service "api-gateway" \
	-proto_path ./../ \
	-idl ./../idl/api/gateway.proto \
	&& go mod tidy

.PHONY: update
update:
	hz update \
	-proto_path ./../ \
	-idl ./../idl/api/gateway.proto

.PHONY: run
run:
	go build -o hertz_test && ./hertz_test
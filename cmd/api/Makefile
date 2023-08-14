include ../common.mk

SERVICE_NAME = api

.PHONY: build
build:
	sh build.sh

.PHONY: new
new:
	hz new \
	-module $(MODULE)/cmd/$(SERVICE_NAME) \
	-service "$(SERVICE_NAME)" \
	-idl ../../idl/api.thrift \

.PHONY: gen
gen:
	hz update \
	-module $(MODULE)/cmd/$(SERVICE_NAME) \
	-idl ../../idl/api.thrift \

.PHONY: server
server:
	make build
	cd output && sh bootstrap.sh
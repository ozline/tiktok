include ../common.mk

SERVICE_NAME = video

.PHONY: server
server:
	sh build.sh
	cd output && sh bootstrap.sh

.PHONY: gen
gen:
	kitex \
	-service "$(SERVICE_NAME)" \
	-module "$(MODULE)" \
	-type protobuf \
	-I ../../idl $(SERVICE_NAME).proto
	go mod tidy

.PHONY: test
client:
	cd test && go test -v

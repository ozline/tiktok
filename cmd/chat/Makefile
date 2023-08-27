include ../common.mk

SERVICE_NAME = chat

.PHONY: server
server:
	sh build.sh
	cd output && sh bootstrap.sh

.PHONY: gen
gen:
	kitex \
	-gen-path ../../kitex_gen \
	-service "$(SERVICE_NAME)" \
	-module "$(MODULE)" \
	-type thrift \
	../../idl/$(SERVICE_NAME).thrift
	go mod tidy

.PHONY: test
test:
	cd ../../test/chat && go test -v
.PHONY: coverage
coverage:
	cd service && go test -coverprofile=../coverage/coverage.out
	go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
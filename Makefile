DIR = $(shell pwd)
CMD = $(DIR)/cmd
CONFIG_PATH = $(DIR)/config
IDL_PATH = $(DIR)/idl
OUTPUT_PATH = $(DIR)/output

SERVICES := api user follow interaction video chat
service = $(word 1, $@)

# mock gen
MOCKS := user_mock
mock = $(word 1, $@)

.PHONY: env-up
env-up:
	docker-compose up -d

.PHONY: env-down
env-down:
	docker-compose down

.PHONY: $(SERVICES)
$(SERVICES):
	mkdir -p output
	cd $(CMD)/$(service) && sh build.sh
	cd $(CMD)/$(service)/output && cp -r . $(OUTPUT_PATH)/$(service)


.PHONY: $(MOCKS)
$(MOCKS):
	@mkdir -p mocks
	mockgen -source=./idl/$(mock).go -destination=./mocks/$(mock).go -package=mocks
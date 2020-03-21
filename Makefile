# ---------- Import make config
conf ?= Makefile.conf
include $(conf)
export $(shell sed 's/=.*//' $(conf))

#
# make <OBJECT> <VERB>
#

# use auto-completion
# complete -W "\`grep -oE '^[a-zA-Z0-9_.-]+:([^=]|$)' Makefile | sed 's/[^a-zA-Z0-9_.-]*$//'\`" make
#

# ---------- Help Menu
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help

# *******************************************
# ***************** DEV *********************
# *******************************************

pb-codegen: ## Generate code out of proto files
	protoc -I ${INCLUDE_PROTO_DIR} ${INPUT_PROTO_FILES} --go_out=plugins=grpc:${PROTO_OUTPUT_DIR}

# ***************** Simple RPC targets *********************

build-client:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${OUTPUT_BIN_DIR}/SimpleClient ./Simple/client/*.go

build-server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${OUTPUT_BIN_DIR}/SimpleServer ./Simple/server/*.go

build-simple: build-client build-server 

simple-container-build: build-simple ## Build containers for simple server and client
	docker build -f SimpleDockerfile -t simple_server:latest .

simple-config-check: ## Verify docker-compose configuration file
	docker-compose -f simple-docker-compose.yml config

simple-start: ## Start docker containers using docker-compose.yml file
	docker-compose -f simple-docker-compose.yml up
	
simple-stop: ## Stop all docker containers started earlier
	docker-compose -f simple-docker-compose.yml down

# ***************** Bidirectional streaming RPC targets *********************

build-streamclient:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${OUTPUT_BIN_DIR}/StreamingServerClient ./Streaming/client/*.go

build-streamserver:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ${GO} build -o ${OUTPUT_BIN_DIR}/StreamingServer ./Streaming/server/*.go

build-streaming: build-streamclient build-streamserver ## Build executable from go code
	
stream-container-build: build-streaming ## Build streaming client and server
	docker build -f StreamDockerfile -t stream_server:latest .

stream-config-check: ## Verify docker-compose configuration file
	docker-compose -f stream-docker-compose.yml config

stream-start: ## Start docker containers using docker-compose.yml file
	docker-compose -f stream-docker-compose.yml up
	
stream-stop: ## Stop all docker containers started earlier
	docker-compose -f stream-docker-compose.yml down

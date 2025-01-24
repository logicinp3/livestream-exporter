APP_NAME = livestream-exporter
APP_VERSION = v1
APP_PORT = 8080
CONTAINER_TOOL := $(shell command -v podman 2> /dev/null || command -v docker 2> /dev/null || echo "none")
CONTAINER_REGISTRY_DOMAIN = harbor.slleisure.com
CONTAINER_REGISTRY_PROJECT = devops
CONTAINER_IMAGE = ${CONTAINER_REGISTRY_DOMAIN}/${CONTAINER_REGISTRY_PROJECT}/${APP_NAME}:${APP_VERSION}

# check container runtime
ifeq ($(CONTAINER_TOOL),none)
    $(error "podman or docker command not found!")
endif


.PHONY: all build push test
all: image clean  ## build, push image and clean

#test: build run clean  ## build, run image and clean
test:
	@echo "Testing..."

build:
	@echo "Building image"
	${CONTAINER_TOOL} build --build-arg APP_PORT=${APP_PORT} -t ${CONTAINER_IMAGE} .

push:
	@echo "Pushing image to registry"
	${CONTAINER_TOOL} push ${CONTAINER_IMAGE}

clean:
	@echo "Cleaning up"
	${CONTAINER_TOOL} image prune || true;\
	${CONTAINER_TOOL} rmi ${CONTAINER_IMAGE} || true

run:
	@echo "Running container"
	${CONTAINER_TOOL} run -p ${APP_PORT}:${APP_PORT} ${CONTAINER_IMAGE}

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

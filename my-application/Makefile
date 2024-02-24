# Build the my-application example

# Usage:
#   [VERSION=latest] [REGISTRY="my-registry"] make build
VERSION?=latest
REGISTRY?=my-registry

release: build push clean

build:
	docker buildx build --load -t ${REGISTRY}/my-application:${VERSION} .

push:
	docker buildx build --push -t ${REGISTRY}/my-application:${VERSION} .

.PHONY: release build push

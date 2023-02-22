SHELL := /bin/bash

run:
	run go run main.go

# =======================================================================
# Building containers

VERSION := 1.0

all: service

service:
	docker build \
		-f zarf/docker/Dockerfile \
		-t service-arm64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.


# =======================================================================
# Running from within k8s/kind

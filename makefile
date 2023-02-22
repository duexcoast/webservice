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

KIND_CLUSTER := ardan-starter-cluster

kind-up:
	kind create cluster \
		--image kindest/node:v1.26.0@sha256:36a22ea7ff0381daf50b12fd1d41e8cd8f6625a4041b6e4d41f084adbe8c2da1 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml 

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide

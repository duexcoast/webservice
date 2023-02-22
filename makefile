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
		--image kindest/node:v1.20.15@sha256:a32bf55309294120616886b5338f95dd98a2f7231519c7dedcec32ba29699394 \
		--name $(KIND_CLUSTER) \
		--config zarf/k8s/kind/kind-config.yaml 

kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

kind-load:
	kind load docker-image service-arm64:$(VERSION) --name $(KIND_CLUSTER)

kind-apply:
	cat zarf/k8s/base/service-pod/base-service.yaml | kubectl apply -f -

kind-logs:
	kubectl logs --namespace=service-system -l app=service --all-containers=true -f --tail=100 

kind-restart:
	kubectl rollout restart deployment service-pod --namespace=service-system

kind-update: all kind-load kind-restart

kind-describe:
	kubectl describe pod -l app=service

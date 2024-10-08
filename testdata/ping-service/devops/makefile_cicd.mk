# build-image
.PHONY: build
# build :-->: service image
build:
	#docker build -t ping-service -f ./testdata/ping-service/devops/Dockerfile .
	#docker pull golang:1.22.8
	#docker pull debian:stable-20240926-slim
	docker build \
		--build-arg BUILD_FROM_IMAGE=golang:1.22.8 \
		--build-arg RUN_SERVICE_IMAGE=debian:stable-20240926-slim \
		--build-arg APP_DIR=testdata \
		--build-arg SERVICE_NAME=ping-service \
		--build-arg VERSION=latest \
		-t ping-service:latest \
		-f ./testdata/ping-service/devops/Dockerfile .


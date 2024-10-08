# build-image
.PHONY: build
# build :-->: service image
build:
	#docker build -t ping-service -f ./testdata/ping-service/devops/Dockerfile .
	docker build \
    		--build-arg APP_DIR=testdata \
    		--build-arg SERVICE_NAME=ping-service \
    		--build-arg VERSION=v1.0.1 \
    		-t ping-service:v1.0.1 \
    		-f ./testdata/ping-service/devops/Dockerfile .


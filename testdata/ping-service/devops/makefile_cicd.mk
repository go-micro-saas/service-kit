# build-image
.PHONY: build
# build :-->: service image
build:
	docker build -t ping-service -f ./devops/Dockerfile .


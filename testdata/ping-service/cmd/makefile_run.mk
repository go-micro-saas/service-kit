# ping-service
.PHONY: run-ping-service
# run service :-->: run ping-service
run-ping-service:
	go run ./testdata/ping-service/cmd/... -conf=./testdata/ping-service/configs

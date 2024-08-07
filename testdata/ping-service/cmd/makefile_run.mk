# ping-service
.PHONY: run-ping-service
# run service :-->: run ping-service
run-ping-service:
	go run ./testdata/ping-service/cmd/... -conf=./testdata/ping-service/configs

.PHONY: testing-ping-service
# testing service :-->: testing ping-service
testing-ping-service:
	curl http://127.0.0.1:10101/api/v1/ping/say_hello && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/logger && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/error && echo "\n"
	curl http://127.0.0.1:10101/api/v1/ping/panic && echo "\n"

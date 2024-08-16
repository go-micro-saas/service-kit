ARG GOPATH=/myworkspace/golang
ARG WORK_DIR=/myworkspace/golang/src/workspace
ARG GOPROXY=https://goproxy.io,direct
ARG GOMODCACHE=/myworkspace/golang/pkg/mod
ARG GOCACHE=/myworkspace/golang/build_cache

# service name
ARG SERVICE_NAME=ping-service

FROM golang:1.22.6 AS builder

# arg
ARG WORK_DIR
ARG GOPATH
ARG GOPROXY
ARG GOMODCACHE
ARG GOCACHE
ARG SERVICE_NAME
ARG VERSION

# env
ENV GOPATH=${GOPATH}
ENV GOPROXY=${GOPROXY}
ENV GOMODCACHE=${GOMODCACHE}
ENV GOCACHE=${GOCACHE}
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR ${WORK_DIR}

COPY . .

RUN go mod download
RUN go mod tidy

# go build -ldflags="-s -w -X 'main.Version=1.0.0' -extldflags '-static'"
RUN go build go build -ldflags="-s -w -X 'main.Version=${VERSION}' -extldflags '-static'" \
    -o ./bin/main \
    ./app/${SERVICE_NAME}/cmd/${SERVICE_NAME}/main.go

FROM debian:stable-slim

# arg
ARG WORK_DIR
ARG SERVICE_NAME

# env
ENV TZ=Asia/Shanghai

WORKDIR ${WORK_DIR}

RUN apt-get update && \
    apt-get install -y --no-install-recommends -y tzdata ca-certificates netbase telnet curl  less && \
    ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone && \
    apt-get autoremove -y && apt-get autoclean -y && apt-get clean -y && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder ${WORK_DIR}/bin/main ./main
COPY --from=builder ${WORK_DIR}/app/${SERVICE_NAME}/configs ./configs

# http
EXPOSE 10001
# grpc
EXPOSE 10002

ENTRYPOINT ./main -conf=./configs

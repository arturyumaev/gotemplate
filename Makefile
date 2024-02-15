PROJECT?=github.com/arturyumaev/gotemplate
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

all: build_api

build_api:
	go build \
		-ldflags '-w -s \
		-X ${PROJECT}/version.Version=${VERSION} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}' \
		-o bin/api cmd/api/main.go 

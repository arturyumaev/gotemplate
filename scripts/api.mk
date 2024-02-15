PROJECT?=github.com/arturyumaev/gotemplate
NAME?=api
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

api_build: clean
	go build \
		-ldflags '-w -s \
		-X ${PROJECT}/version.APIVersion=${VERSION} \
		-X ${PROJECT}/version.APIName=${NAME} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}' \
		-o bin/api cmd/api/main.go 

api_start: api_build
	./bin/api

api_status:
	curl localhost:3000/status | jq

PROJECT?=github.com/arturyumaev/gotemplate
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

api_build: clean
	go build \
		-ldflags '-w -s \
		-X ${PROJECT}/version.APIVersion=${VERSION} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}' \
		-o bin/api cmd/api/main.go 

api_run: api_build
	./bin/api

api_status:
	curl localhost:3000/status | jq

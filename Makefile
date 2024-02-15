PROJECT?=github.com/arturyumaev/gotemplate
API_VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

all: api_run

clean:
	rm -rf bin

api_build: clean
	go build \
		-ldflags '-w -s \
		-X ${PROJECT}/version.APIVersion=${API_VERSION} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}' \
		-o bin/api cmd/api/main.go 

api_run: api_build
	./bin/api

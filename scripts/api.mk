PROJECT?=github.com/arturyumaev/gotemplate
NAME?=api
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')

clean:
	rm bin/api

api_build: clean
	go build \
		-ldflags "-w -s \
		-X ${PROJECT}/version.APIVersion=${VERSION} \
		-X ${PROJECT}/version.APIName=${NAME} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o bin/api cmd/api/main.go 

api_build_docker:
	docker build -t ${NAME}:${VERSION} -t ${NAME}:latest -f deployments/production/Dockerfile.api .

api_start: api_build
	./bin/api

api_start_docker: api_build_docker
	docker run --rm -p 3000:3000 --env-file deployments/production/.env ${NAME}:${VERSION}

api_status:
	curl localhost:3000/status | jq

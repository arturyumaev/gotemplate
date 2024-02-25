PROJECT?=github.com/arturyumaev/gotemplate
NAME?=api
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date '+%Y-%m-%dT%H:%M:%S')
CONTAINER_NAME?=docker.io/arturyumaev/${NAME}

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

api_start: api_build
	./bin/api

api_docker_build:
	docker build -t ${CONTAINER_NAME}:${VERSION} -t ${CONTAINER_NAME}:latest -f deployments/production/Dockerfile.api .

api_docker_start: api_docker_build
	docker run --rm -p 3000:3000 --env-file deployments/production/.env ${CONTAINER_NAME}:${VERSION}

api_docker_push:
	docker push ${CONTAINER_NAME}:${VERSION}
	docker push ${CONTAINER_NAME}:latest

api_status:
	curl localhost:3000/status | jq

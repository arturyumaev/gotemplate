include scripts/*.mk

dev-up: api_docker_build
	docker compose -v -f deployments/development/docker-compose.yaml up -d

dev-down:
	docker compose -f deployments/development/docker-compose.yaml down

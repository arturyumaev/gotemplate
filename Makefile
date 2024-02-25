include scripts/*.mk

dev-up: api_build_docker
	docker compose -f deployments/development/docker-compose.yaml up -d

dev-down:
	docker compose -f deployments/development/docker-compose.yaml down

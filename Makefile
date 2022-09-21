docker:
	DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up

docker-build:
	DOCKER_BUILDKIT=1 COMPOSE_DOCKER_CLI_BUILD=1 docker-compose up --build

docker-down:
	docker-compose down && docker-compose rm

db:
	docker-compose up postgres

local:
	SPEAKEASY_SERVER_URL=localhost:9090 SPEAKEASY_ENVIRONMENT=local POSTGRES_DSN=postgres://guest:guest@localhost:5431/speakeasy?sslmode=disable SPEAKEASY_SDK_API_KEY=$${api_key:?please set \$$api_key} go run cmd/server/main.go

healthcheck:
	curl -vv http://localhost:8081/health
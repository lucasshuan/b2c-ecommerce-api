gqlgen:
	@echo "Running go generate..."
	go run github.com/99designs/gqlgen generate

server:
	@echo "Running server..."
	go run cmd/main.go

docker-up:
	@echo "Running Docker containers..."
	@for compose_dir in ./docker/*/compose.yml; do \
		project_name=$$(basename $$(dirname "$$compose_dir")); \
		docker-compose -p $$project_name -f "$$compose_dir" up -d; \
	done

docker-down:
	@echo "Stopping Docker containers..."
	@for compose_dir in ./docker/*/compose.yml; do \
		project_name=$$(basename $$(dirname "$$compose_dir")); \
		docker-compose -p $$project_name -f "$$compose_dir" down -v; \
	done

deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod vendor
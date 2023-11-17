server:
	@echo "Running server..."
	go run cmd/main.go

compose:
	@echo "Running Docker containers..."
	@for compose_dir in ./docker/*/compose.yml; do \
		project_name=$$(basename $$(dirname "$$compose_dir")); \
		docker-compose -p $$project_name -f "$$compose_dir" up -d; \
	done
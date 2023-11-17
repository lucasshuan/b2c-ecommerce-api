dev:
	@echo "Running server locally..."
	go run cmd/main.go

prod:
	@echo "Running server in production..."
	go run cmd/main.go

docker:
	@echo "Running Docker containers..."
	@for compose_file in ./configs/docker/*.yml; do \
		docker-compose -p ecommerce-api -f "$$compose_file" up -d; \
	done

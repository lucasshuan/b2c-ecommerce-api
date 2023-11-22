gqlgen:
	go run github.com/99designs/gqlgen

server:
	go run main.go server

docker-up:
	go run main.go docker-up

docker-down:
	go run main.go docker-down

deps:
	go mod tidy
	go mod vendor
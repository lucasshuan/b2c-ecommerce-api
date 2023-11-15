package main

import (
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/server"
)

func main() {
	database.Initialize()
	server.Start()
}

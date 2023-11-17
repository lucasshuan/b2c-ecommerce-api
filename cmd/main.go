package main

import (
	"github.com/nobeluc/ecommerce-api/configs"
	"github.com/nobeluc/ecommerce-api/internal/database"
	"github.com/nobeluc/ecommerce-api/internal/log"
	"github.com/nobeluc/ecommerce-api/internal/server"
)

func main() {
	log.Init()

	c := configs.Load()

	database.Init(c)
	server.Init(c)
}

package main

import (
	"log"

	"github.com/savanyv/e-commerce/config"
	"github.com/savanyv/e-commerce/internal/app"
)

func main() {
	config := config.LoadConfig()

	server := app.NewServer(config)
	if err := server.RunServer(); err != nil {
		log.Println("error running server", err)
	}
}

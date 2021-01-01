package main

import (
	"github.com/joho/godotenv"
	"github.com/luschnat-ziegler/cc_backend_go/app"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app.Start()
}

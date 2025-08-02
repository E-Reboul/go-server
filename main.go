package main

import (
	"github.com/E-Reboul/go-server/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server.Launch()
}

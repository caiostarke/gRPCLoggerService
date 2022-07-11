package main

import (
	"github.com/caiostarke/apiAndGRPC/logger_service/app/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	server.RunGRPCServer()
}

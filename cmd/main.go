package main

import (
	"os"

	"github.com/Imanr2/Restaurant_API/internal/application"
	"github.com/Imanr2/Restaurant_API/internal/database"
	"github.com/joho/godotenv"
)

const (
	DefaultEnvFilePath = ".env"

	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NET      = "DB_NET"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
	JWT_KEY     = "JWT_KEY"
)

var jwtKey = []byte(os.Getenv(JWT_KEY))

func main() {
	godotenv.Load("../.env")

	dbConfig := database.NewDBConfig(
		os.Getenv(DB_USER),
		os.Getenv(DB_PASSWORD),
		os.Getenv(DB_NET),
		os.Getenv(DB_PORT),
		os.Getenv(DB_NAME),
	)

	a := application.Application{}
	a.Initialize(*dbConfig)

	a.Run()
}

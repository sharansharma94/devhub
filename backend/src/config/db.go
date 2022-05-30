package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
}

func GetDBConfig() Config {

	var err error
	if os.Getenv("GIN_MODE") == "release" {
		err = godotenv.Load(".env-prod")
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatalf("Error occured while loading env file %s ", err)
	}

	config := Config{POSTGRES_USER: os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}

	fmt.Println(config)
	return config
}

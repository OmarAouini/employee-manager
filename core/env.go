package core

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//only for local development, in prod is ignored
func LoadDotEnv() {
	fmt.Println("\nconfigure env variables...")
	if os.Getenv("APP_ENV") == "local" {
		fmt.Println("environment local")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

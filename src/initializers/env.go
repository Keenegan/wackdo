package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		// In production (Render), env vars are injected directly, no .env file needed
		log.Println("No .env file found, using environment variables")
	}
}

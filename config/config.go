package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(ENV_VAR string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load environment variables")
	}
	envVar := os.Getenv(ENV_VAR)
	if envVar == "" {
		log.Fatal("openai api key is not provided.")
	}
	return envVar
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load environment variables")
	}

	OpenAI := os.Getenv("OPENAI_KEY")
	if OpenAI == "" {
		log.Fatal("openai api key is not provided.")
	}

	return OpenAI
}

func GetCMCKey() string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load CMC API key")
	}

	cmc_api_key := os.Getenv("CMC_PRO_API_KEY")
	if cmc_api_key == "" {
		log.Fatal("cmc api key is not provided")
	}

	return cmc_api_key
}

func GetRapidAPI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load rapid API key.")
	}

	rapidAPIKey := os.Getenv("RAPIDAPI_KEY")
	if rapidAPIKey == "" {
		log.Fatal("Please provide RAPIDAPI key to .env")
	}
	return rapidAPIKey
}

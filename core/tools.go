package core

import config "github.com/siddhantprateek/go-gpt/config"

var (
	API_KEY_TITLE      = "X-RapidAPI-Key"
	API_KEY            = config.GetEnv("API_KEY")
	API_KEY_HOST_TITLE = "X-RapidAPI-Host"
	API_KEY_HOST       = "openai80.p.rapidapi.com"
)

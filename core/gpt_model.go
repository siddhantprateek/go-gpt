package core

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhantprateek/go-gpt/config"
	"github.com/siddhantprateek/go-gpt/model"
)

func GPTModels() []string {

	model_url := "https://openai80.p.rapidapi.com/models"

	req, err := http.NewRequest("GET", model_url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-RapidAPI-Key", config.GetRapidAPI())
	req.Header.Add("X-RapidAPI-Host", "openai80.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// Decode the response body into a GPTModelRespone struct
	var modelReponse model.GPTModelResponse
	err = json.NewDecoder(res.Body).Decode(&modelReponse)
	if err != nil {
		log.Fatal("Unable to decode the GPTModel Response.")
	}

	var modelList []string
	for _, ModelObj := range modelReponse.Data {
		modelList = append(modelList, ModelObj.Root)
	}
	return modelList
}

package core

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhantprateek/go-gpt/model"
)

func GetAllModels() []string {
	requestURL := "https://openai80.p.rapidapi.com/models"
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add(API_KEY_TITLE, API_KEY)
	req.Header.Add(API_KEY_HOST_TITLE, API_KEY_HOST)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

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

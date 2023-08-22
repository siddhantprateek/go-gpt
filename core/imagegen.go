package core

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/siddhantprateek/go-gpt/model"
)

func ImageGeneration(image_desp string) (string, string) {
	requestURL := "https://openai80.p.rapidapi.com/images/generations"
	reqBody := model.GPTImageReqModel{
		Prompt: image_desp,
		N:      2,
		Size:   "",
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Unable to Marshall Request Body.")
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
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

	var imageResponse model.GPTImageModel
	err = json.NewDecoder(res.Body).Decode(&imageResponse)
	if err != nil {
		log.Fatal("Unable to decode the Image Response.")
	}

	return imageResponse.Data[0].URL, imageResponse.Data[1].URL
}

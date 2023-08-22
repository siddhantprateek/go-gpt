package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/siddhantprateek/go-gpt/model"
)

func TxtModeration(input string) {
	requestURL := "https://openai80.p.rapidapi.com/moderations"
	reqPayload := model.GPTModerationModel{
		Input: input,
		Model: "text-moderation-stable",
	}

	payload, err := json.Marshal(reqPayload)
	if err != nil {
		log.Fatal("Unable to Marshall the request body.")
	}

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal("unable to make txt moderation request.")
	}
	req.Header.Add(API_KEY_TITLE, API_KEY)
	req.Header.Add(API_KEY_HOST_TITLE, API_KEY_HOST)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("unable to recieve txt moderation response.")
	}

	defer res.Body.Close()

	var moderationResponse model.GPTModerationModelResponse
	err = json.NewDecoder(res.Body).Decode(&moderationResponse)
	if err != nil {
		log.Fatal("unable to Decode the response.")
	}
	fmt.Println("---------- : Category Score : ----------- ")
	fmt.Println("Hate: ", int(moderationResponse.Results[0].CategoryScores.Hate))

}

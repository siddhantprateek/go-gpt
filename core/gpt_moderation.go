package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "github.com/siddhantprateek/go-gpt/config"
	"github.com/siddhantprateek/go-gpt/model"
)

func GPTModeration(input string) {

	openAI_url := "https://openai80.p.rapidapi.com/moderations"

	reqBody := model.GPTModerationModel{
		Input: input,
		Model: "text-moderation-stable",
	}
	fmt.Println("resquest input: ", input)
	payload, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Unable to Marshall the request body.")
	}

	req, err := http.NewRequest("POST", openAI_url, bytes.NewBuffer(payload))
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

	var moderationResponse model.GPTModerationModelResponse
	err = json.NewDecoder(res.Body).Decode(&moderationResponse)
	if err != nil {
		log.Fatal("Unable to Decode the response.")
	}
	fmt.Println("---------- : Category Score : ----------- ")
	fmt.Println("Hate: ", int(moderationResponse.Results[0].CategoryScores.Hate))
	// fmt.Println("HateThreatening: ", moderationResponse.Results[0].CategoryScores.HateThreatening)
	// fmt.Println("SelfHarm: ", moderationResponse.Results[0].CategoryScores.SelfHarm)
	// fmt.Println("Sexual: ", moderationResponse.Results[0].CategoryScores.Sexual)
	// fmt.Println("SexualMinors: ", moderationResponse.Results[0].CategoryScores.SexualMinors)
	// fmt.Println("Violence: ", moderationResponse.Results[0].CategoryScores.Violence)
	// fmt.Println("ViolenceGraphic: ", moderationResponse.Results[0].CategoryScores.ViolenceGraphic)

}

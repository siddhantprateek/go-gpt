package core

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/siddhantprateek/go-gpt/model"
)

func SentimentAnalysis(input string) model.SentimentPayload {
	requestURL := "https://open-ai21.p.rapidapi.com/sentiment"
	payload := strings.NewReader("text=" + input)
	req, err := http.NewRequest("POST", requestURL, payload)
	if err != nil {
		log.Fatal("Unable to create the request.")
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", "befc0ba1e5mshfe351ad61d3b7e1p177691jsn9cca0b5511ab")
	req.Header.Add("X-RapidAPI-Host", "open-ai21.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
	}

	var sentimentResp model.SentimentPayload
	err = json.NewDecoder(bytes.NewReader(resBody)).Decode(&sentimentResp)
	if err != nil {
		log.Fatal("Unable to decode the body.")
	}

	return sentimentResp
}

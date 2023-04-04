package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	config "github.com/siddhantprateek/go-gpt/config"
	models "github.com/siddhantprateek/go-gpt/model"
)

func RapidChatGPT(content string) string {

	rapid_url := "https://openai80.p.rapidapi.com/chat/completions"

	model := "gpt-3.5-turbo"
	var msg []models.Messages
	msg = append(msg, models.Messages{
		Role:    "user",
		Content: content,
	})
	requestData := models.ChatCompletionReqModel{
		Model:    model,
		Messages: msg,
	}
	payload, err := json.Marshal(requestData)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", rapid_url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", config.GetRapidAPI())
	req.Header.Add("X-RapidAPI-Host", "openai80.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("response.json")
	if err != nil {
		log.Fatal("Unable to create the file.")
	}

	defer res.Body.Close()
	body, _ := io.Copy(outputFile, res.Body)
	fmt.Println(string(rune(body)))
	responseFile, err := os.Open("response.json")
	if err != nil {
		log.Fatal("Unable to open response file.")
	}

	defer responseFile.Close()

	decoder := json.NewDecoder(responseFile)

	var chatResponse models.ChatCompletionRespModel
	err = decoder.Decode(&chatResponse)
	if err != nil {
		log.Fatal(err)
	}

	// responseFile, err := ioutil.ReadFile("response.json")
	// if err != nil {
	// 	panic(err)
	// }

	// var chatResponse models.ChatCompletionRespModel
	// err = json.Unmarshal([]byte(responseFile), &chatResponse)
	// if err != nil {
	// 	panic(err)
	// }
	// content := completion.Choices[0].Message.Content
	responseContent := chatResponse.Choices[0].Message.Content
	return responseContent

}

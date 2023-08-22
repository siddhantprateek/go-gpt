package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	models "github.com/siddhantprateek/go-gpt/model"
)

// type GPTInstance struct {
// 	prompt_token      int
// 	completion_tokens int
// 	total_tokens      int
// }

func RapidChatGPT(content string) string {
	requestURL := "https://openai80.p.rapidapi.com/chat/completions"
	// set model
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

	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add(API_KEY_TITLE, API_KEY)
	req.Header.Add(API_KEY_HOST_TITLE, API_KEY_HOST)

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

	var chatResponse models.ChatCompletionRespModel
	err = json.NewDecoder(responseFile).Decode(&chatResponse)
	if err != nil {
		log.Fatal(err)
	}

	responseContent := chatResponse.Choices[0].Message.Content
	return responseContent
}

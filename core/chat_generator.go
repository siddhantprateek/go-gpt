package core

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	config "github.com/siddhantprateek/go-gpt/config"
)

func ChatGPT(chat_message string) string {

	// Provide the OpenAI API key
	openai_apiKey := config.GetEnv()
	client := openai.NewClient(openai_apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: chat_message,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("Chat Completion Error: %v\n", err)
	}

	// fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}

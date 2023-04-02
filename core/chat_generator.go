package core

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func ChatGPT() string {

	// Provide the OpenAI API key
	client := openai.NewClient("")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello",
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

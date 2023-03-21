package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Responder interface {
	Respond(string) (Response, error)
}
type openAIConn struct {
	client *openai.Client
}

type Response struct {
	ID      string
	Message string
	Time    int64
	Choices int
}

func (conn openAIConn) Respond(input string) (Response, error) {
	// cli := openai.NewClient("")

	response, err := conn.client.CreateChatCompletion(context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a chat bot",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		return Response{}, fmt.Errorf("error requesting chat completion from open API: %q", err)

	}
	return Response{
		ID:      response.ID,
		Message: response.Choices[0].Message.Content,
		Time:    response.Created,
		Choices: len(response.Choices),
	}, nil
}

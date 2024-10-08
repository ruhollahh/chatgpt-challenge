package openaiclient

import (
	"chatgpt-challenge/internal/param/openai"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type Config struct {
	AuthToken            string
	Model                string
	MaxCompletionsTokens int
}

type Adapter struct {
	cfg    Config
	client *openai.Client
}

func New(cfg Config) Adapter {
	client := openai.NewClient(cfg.AuthToken)
	return Adapter{cfg: cfg, client: client}
}

func (a Adapter) StructuredOutput(req openaiparam.StructuredOutputRequest) (openaiparam.StructuredOutputResponse, error) {
	resp, err := a.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:                a.cfg.Model,
			MaxCompletionsTokens: a.cfg.MaxCompletionsTokens,
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:        req.JSONSchema.Name,
					Description: req.JSONSchema.Description,
					Schema:      req.JSONSchema.Schema,
					Strict:      true,
				},
			},
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: req.SystemMessage,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: req.UnstructuredText,
				},
			},
		},
	)

	if err != nil {
		return openaiparam.StructuredOutputResponse{}, fmt.Errorf("error creating chat completion: %s", err.Error())
	}

	return openaiparam.StructuredOutputResponse{
		StructuredJSON: resp.Choices[0].Message.Content,
	}, nil
}

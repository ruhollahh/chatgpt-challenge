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

// Client is a wrapper around the OpenAI client that uses a configuration struct to store
// necessary settings. It provides methods for interacting with OpenAI's APIs.
type Client struct {
	cfg Config
	c   *openai.Client
}

func New(cfg Config) Client {
	client := openai.NewClient(cfg.AuthToken)
	return Client{cfg: cfg, c: client}
}

// StructuredOutput sends a request to the OpenAI API to transform unstructured text into structured JSON data
// based on the schema provided in the request.
//
// The function uses the ChatCompletion API with the specified model and token limits from the Client configuration.
// It formats the response as JSON according to the provided schema.
func (a Client) StructuredOutput(req openaiparam.StructuredOutputRequest) (openaiparam.StructuredOutputResponse, error) {
	resp, err := a.c.CreateChatCompletion(
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

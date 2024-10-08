package laptopstructify

import (
	"chatgpt-challenge/internal/param/openai"
)

type Config struct {
	SystemMessage string
}

type StructGenerator interface {
	StructuredOutput(request openaiparam.StructuredOutputRequest) (openaiparam.StructuredOutputResponse, error)
}

type Service struct {
	cfg             Config
	structGenerator StructGenerator
}

func New(cfg Config, s StructGenerator) Service {
	return Service{
		cfg:             cfg,
		structGenerator: s,
	}
}

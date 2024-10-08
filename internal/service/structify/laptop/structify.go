package laptopstructify

import (
	openaiparam "chatgpt-challenge/internal/param/openai"
	laptopstructifyparam "chatgpt-challenge/internal/param/structify/laptop"
	"chatgpt-challenge/internal/schema/laptop"
	"fmt"
)

func (s Service) Structify(req laptopstructifyparam.StructifyRequest) (laptopstructifyparam.StructifyResponse, error) {
	laptopSchema := laptopschema.New()
	res, err := s.structGenerator.StructuredOutput(openaiparam.StructuredOutputRequest{
		UnstructuredText: req.Content,
		JSONSchema: openaiparam.JSONSchema{
			Name:        laptopSchema.Name(),
			Description: laptopSchema.Description(),
			Schema:      laptopSchema,
		},
		SystemMessage: s.cfg.SystemMessage,
	})
	if err != nil {
		return laptopstructifyparam.StructifyResponse{}, fmt.Errorf("error generating laptop structure: %s", err.Error())
	}

	laptop, err := laptopSchema.ParseContent(res.StructuredJSON)
	if err != nil {
		return laptopstructifyparam.StructifyResponse{}, fmt.Errorf("error parsing generated structure: %s", err.Error())
	}

	return laptopstructifyparam.StructifyResponse{Laptop: laptop}, nil
}

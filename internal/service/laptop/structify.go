package laptopservice

import (
	"chatgpt-challenge/internal/entity"
	laptopparam "chatgpt-challenge/internal/param/laptop"
	openaiparam "chatgpt-challenge/internal/param/openai"
	"fmt"
)

// Structify processes an unstructured laptop description and converts it into a structured JSON format.
// It uses a predefined JSON schema to validate the structure of the laptop data.
//
// The function first generates a structured output using the OpenAI API, then parses the structured JSON
// into a laptop entity. If the content is not valid according to the schema, an error is returned.
func (s Service) Structify(req laptopparam.StructifyRequest) (laptopparam.StructifyResponse, error) {
	res, err := s.structGenerator.StructuredOutput(openaiparam.StructuredOutputRequest{
		UnstructuredText: req.Content,
		JSONSchema: openaiparam.JSONSchema{
			Name:        s.laptopSchema.Name(),
			Description: s.laptopSchema.Description(),
			Schema:      s.laptopSchema,
		},
		SystemMessage: s.cfg.SystemMessage,
	})
	if err != nil {
		return laptopparam.StructifyResponse{}, fmt.Errorf("error generating laptop structure: %s", err.Error())
	}

	laptop, err := s.laptopSchema.ParseContent(res.StructuredJSON)
	if err != nil {
		return laptopparam.StructifyResponse{}, fmt.Errorf("error parsing generated structure: %s", err.Error())
	}

	if isValid := ValidateLaptop(laptop); !isValid {
		return laptopparam.StructifyResponse{}, fmt.Errorf("invalid laptop content: %s", res.StructuredJSON)
	}

	return laptopparam.StructifyResponse{Laptop: laptop}, nil
}

func ValidateLaptop(l entity.Laptop) bool {
	if l.Brand == "" {
		return false
	}
	if l.Model == "" {
		return false
	}
	if l.Processor == "" {
		return false
	}
	if l.RamCapacity == "" {
		return false
	}
	if l.RamType == "" {
		return false
	}
	if l.StorageCapacity == "" {
		return false
	}
	if l.BatteryStatus == "" {
		return false
	}

	return true
}

package laptopschema

import (
	"chatgpt-challenge/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai/jsonschema"
	"log"
)

// Schema represents the JSON schema for a Laptop entity.
// It holds the schema's name, description, and definition, which is used for
// marshaling and unmarshaling laptop data.
type Schema struct {
	name        string
	description string
	definition  *jsonschema.Definition
}

// New generates a JSON schema definition based on the Laptop entity structure.
// If schema generation fails, it logs the error and terminates the application.
func New() Schema {
	definition, err := jsonschema.GenerateSchemaForType(&entity.Laptop{})
	if err != nil {
		log.Fatalf("error generating laptop schema: %s", err.Error())
	}

	bsDef := definition.Properties["battery_status"]
	bsDef.Enum = []string{"Yes", "No"}
	definition.Properties["battery_status"] = bsDef

	return Schema{
		name:        "Laptop",
		description: "Laptop Information",
		definition:  definition,
	}
}

// MarshalJSON outputs the JSON serialized version of the schema.
// we send this to the OpenAI's structured output API.
func (s Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.definition)
}

// ParseContent is used to unmarshal the JSON output of the OpenAI's API into a Laptop entity.
func (s Schema) ParseContent(content string) (entity.Laptop, error) {
	laptop := entity.Laptop{}
	err := s.definition.Unmarshal(content, &laptop)
	if err != nil {
		return entity.Laptop{}, fmt.Errorf("error unmarshaling schema: %s", err.Error())
	}

	return laptop, nil
}

func (s Schema) Name() string {
	return s.name
}

func (s Schema) Description() string {
	return s.description
}

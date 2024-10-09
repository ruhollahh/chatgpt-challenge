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

// New creates a new Schema instance for the Laptop entity.
// It generates a JSON schema definition based on the Laptop entity structure.
// The method also sets the allowed values for the "battery_status" field to "Yes" or "No".
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

// MarshalJSON marshals the Schema definition into a JSON format.
// The output of this function is what we send to the OpenAI API
// to establish a schema for it's structured output.
func (s Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.definition)
}

// ParseContent is used to unmarshal the JSON output of the OpenAI API into a Laptop entity.
// It uses the Schema's definition to validate and extract fields from the content.
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

package laptopschema

import (
	"chatgpt-challenge/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai/jsonschema"
	"log"
)

type Schema struct {
	name        string
	description string
	definition  *jsonschema.Definition
}

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

func (s Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.definition)
}

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

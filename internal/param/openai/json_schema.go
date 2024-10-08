package openaiparam

import "encoding/json"

type JSONSchema struct {
	Name        string
	Description string
	Schema      json.Marshaler
}

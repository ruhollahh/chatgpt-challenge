package openaiparam

type StructuredOutputRequest struct {
	UnstructuredText string
	JSONSchema       JSONSchema
	SystemMessage    string
}

type StructuredOutputResponse struct {
	StructuredJSON string
}

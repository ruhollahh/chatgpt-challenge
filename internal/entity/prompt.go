package entity

type PromptStatus string

const (
	PromptStatusPending   PromptStatus = "PENDING"
	PromptStatusFailed    PromptStatus = "FAILED"
	PromptStatusProcessed PromptStatus = "PROCESSED"
)

type Prompt struct {
	ID           string
	Content      string
	Status       PromptStatus
	ErrorMessage string
}

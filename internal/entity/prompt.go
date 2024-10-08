package entity

type PromptStatus string

const (
	PromptStatusPending   PromptStatus = "PENDING"
	PromptStatusProcessed PromptStatus = "PROCESSED"
)

type Prompt struct {
	ID      string
	Content string
	Status  PromptStatus
}

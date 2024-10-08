package entity

type PromptID [32]byte

type PromptStatus string

const (
	PromptStatusPending   PromptStatus = "PENDING"
	PromptStatusProcessed PromptStatus = "PROCESSED"
)

type Prompt struct {
	ID      PromptID
	Content string
	Status  PromptStatus
}

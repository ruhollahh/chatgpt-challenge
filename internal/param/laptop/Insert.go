package laptopparam

import "chatgpt-challenge/internal/entity"

type InsertRequest struct {
	PromptID string
	Laptop   entity.Laptop
}

package promptparam

import "chatgpt-challenge/internal/entity"

type GetAllResponse struct {
	ID           string              `json:"id"`
	Content      string              `json:"content"`
	Status       entity.PromptStatus `json:"status"`
	ErrorMessage string              `json:"error_message"`
}

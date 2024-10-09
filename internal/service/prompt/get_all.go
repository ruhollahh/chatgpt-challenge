package promptservice

import (
	promptparam "chatgpt-challenge/internal/param/prompt"
)

func (s Service) GetAll() []promptparam.GetAllResponse {
	prompts := s.promptRepo.GetAll()
	response := make([]promptparam.GetAllResponse, len(prompts))
	for i, prompt := range prompts {
		response[i] = promptparam.GetAllResponse{
			ID:           prompt.ID,
			Content:      prompt.Content,
			Status:       prompt.Status,
			ErrorMessage: prompt.ErrorMessage,
		}
	}

	return response
}

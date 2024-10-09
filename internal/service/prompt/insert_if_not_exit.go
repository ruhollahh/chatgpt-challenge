package promptservice

import (
	"chatgpt-challenge/internal/entity"
	promptparam "chatgpt-challenge/internal/param/prompt"
)

func (s Service) InsertIfNotExist(req promptparam.InsertIfNotExistRequest) promptparam.InsertIfNotExistResponse {
	inserted := s.promptRepo.SetNX(entity.Prompt{
		ID:           req.ID,
		Content:      req.Content,
		Status:       entity.PromptStatusPending,
		ErrorMessage: "",
	})

	return promptparam.InsertIfNotExistResponse{
		Inserted: inserted,
	}
}

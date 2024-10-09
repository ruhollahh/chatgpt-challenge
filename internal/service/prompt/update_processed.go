package promptservice

import (
	"chatgpt-challenge/internal/entity"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"fmt"
)

func (s Service) UpdateProcessed(req promptparam.UpdateProcessedRequest) error {
	err := s.promptRepo.UpdateStatus(req.ID, entity.PromptStatusProcessed)
	if err != nil {
		return fmt.Errorf("error updating processed status: %w", err)
	}

	return nil
}

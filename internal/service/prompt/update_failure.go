package promptservice

import (
	"chatgpt-challenge/internal/entity"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"fmt"
)

func (s Service) UpdateFailure(req promptparam.UpdateFailureRequest) error {
	err := s.promptRepo.UpdateStatus(req.ID, entity.PromptStatusFailed)
	if err != nil {
		return fmt.Errorf("error updating status: %w", err)
	}

	err = s.promptRepo.UpdateErrorMessage(req.ID, req.ErrorMessage)
	if err != nil {
		return fmt.Errorf("error updating error message: %w", err)
	}

	return nil
}

package promptservice

import "chatgpt-challenge/internal/entity"

//go:generate mockery --name PromptRepo
type PromptRepo interface {
	SetNX(entity.Prompt) bool
	GetAll() []entity.Prompt
	UpdateStatus(string, entity.PromptStatus) error
	UpdateErrorMessage(id string, errMessage string) error
}

type Service struct {
	promptRepo PromptRepo
}

func New(promptRepo PromptRepo) Service {
	return Service{
		promptRepo: promptRepo,
	}
}

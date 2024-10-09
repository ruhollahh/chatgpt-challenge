package prompthandler

import (
	promptservice "chatgpt-challenge/internal/service/prompt"
)

type Handler struct {
	promptSvc promptservice.Service
}

func New(promptSvc promptservice.Service) Handler {
	return Handler{
		promptSvc: promptSvc,
	}
}

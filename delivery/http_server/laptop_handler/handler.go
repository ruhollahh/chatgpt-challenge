package laptophandler

import laptopservice "chatgpt-challenge/internal/service/laptop"

type Handler struct {
	laptopSvc laptopservice.Service
}

func New(laptopSvc laptopservice.Service) Handler {
	return Handler{
		laptopSvc: laptopSvc,
	}
}

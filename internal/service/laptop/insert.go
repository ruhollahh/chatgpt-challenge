package laptopservice

import laptopparam "chatgpt-challenge/internal/param/laptop"

func (s Service) Insert(req laptopparam.InsertRequest) {
	s.laptopRepo.Set(req.PromptID, req.Laptop)
}

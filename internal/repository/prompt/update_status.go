package promptrepo

import (
	"chatgpt-challenge/internal/entity"
	"chatgpt-challenge/internal/repository"
)

func (r *Repository) UpdateStatus(id string, status entity.PromptStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	prompt, ok := r.db[id]
	if !ok {
		return repository.ErrNotFound
	}
	prompt.Status = status
	r.db[id] = prompt

	return nil
}

package promptrepo

import (
	"chatgpt-challenge/internal/repository"
)

func (r *Repository) UpdateErrorMessage(id string, errMessage string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	prompt, ok := r.db[id]
	if !ok {
		return repository.ErrNotFound
	}
	prompt.ErrorMessage = errMessage
	r.db[id] = prompt

	return nil
}

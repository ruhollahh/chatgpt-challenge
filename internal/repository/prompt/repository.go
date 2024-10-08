package promptrepo

import (
	"chatgpt-challenge/internal/entity"
	"chatgpt-challenge/internal/repository"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db map[string]entity.Prompt
}

func New() *Repository {
	return &Repository{
		db: make(map[string]entity.Prompt),
	}
}

func (r *Repository) GetAll() []entity.Prompt {
	r.mu.RLock()
	defer r.mu.RUnlock()
	prompts := make([]entity.Prompt, 0, len(r.db))
	for _, p := range r.db {
		prompts = append(prompts, p)
	}

	return prompts
}

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

// SetNX sets only if the prompt does not already exist
// return true if set and false if not set
func (r *Repository) SetNX(prompt entity.Prompt) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.db[prompt.ID]; ok {
		return false
	}
	r.db[prompt.ID] = prompt
	return true
}

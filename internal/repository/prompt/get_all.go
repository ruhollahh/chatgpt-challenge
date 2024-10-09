package promptrepo

import "chatgpt-challenge/internal/entity"

func (r *Repository) GetAll() []entity.Prompt {
	r.mu.RLock()
	defer r.mu.RUnlock()
	prompts := make([]entity.Prompt, 0, len(r.db))
	for _, p := range r.db {
		prompts = append(prompts, p)
	}

	return prompts
}

package promptrepo

import "chatgpt-challenge/internal/entity"

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

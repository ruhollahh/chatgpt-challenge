package promptrepo

import "chatgpt-challenge/internal/entity"

// SetNX attempts to set a new prompt in the repository if the key (prompt.ID) does not already exist.
// Returns true if the prompt was successfully set (i.e., the key did not previously exist).
// Returns false if the prompt ID already exists in the repository.
func (r *Repository) SetNX(prompt entity.Prompt) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.db[prompt.ID]; ok {
		return false
	}
	r.db[prompt.ID] = prompt
	return true
}

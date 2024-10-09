package laptoprepo

import "chatgpt-challenge/internal/entity"

func (r *Repository) Set(promptID string, laptop entity.Laptop) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.db[promptID] = laptop
}

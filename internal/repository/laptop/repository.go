package laptoprepo

import (
	"chatgpt-challenge/internal/entity"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db map[entity.PromptID]entity.Laptop
}

func New() *Repository {
	return &Repository{
		db: make(map[entity.PromptID]entity.Laptop),
	}
}

func (r *Repository) GetAll() []entity.Laptop {
	r.mu.RLock()
	defer r.mu.RUnlock()
	laptops := make([]entity.Laptop, 0, len(r.db))
	for _, laptop := range r.db {
		laptops = append(laptops, laptop)
	}

	return laptops
}

func (r *Repository) Set(promptID entity.PromptID, laptop entity.Laptop) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.db[promptID] = laptop
}

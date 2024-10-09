package laptoprepo

import (
	"chatgpt-challenge/internal/entity"
	"sync"
)

type Repository struct {
	mu sync.RWMutex
	db map[string]entity.Laptop
}

func New() *Repository {
	return &Repository{
		db: make(map[string]entity.Laptop),
	}
}

package laptoprepo

import "chatgpt-challenge/internal/entity"

func (r *Repository) GetAll() []entity.Laptop {
	r.mu.RLock()
	defer r.mu.RUnlock()
	laptops := make([]entity.Laptop, 0, len(r.db))
	for _, laptop := range r.db {
		laptops = append(laptops, laptop)
	}

	return laptops
}

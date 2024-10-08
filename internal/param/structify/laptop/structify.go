package laptopstructifyparam

import (
	"chatgpt-challenge/internal/entity"
)

type StructifyRequest struct {
	Content string
}

type StructifyResponse struct {
	Laptop entity.Laptop
}

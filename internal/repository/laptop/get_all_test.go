package laptoprepo

import (
	"chatgpt-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Retrieve all laptops when the repository is populated
func TestGetAll(t *testing.T) {
	repo := New()
	laptop1 := entity.Laptop{Brand: "Dell", Model: "Inspiron", Processor: "Intel Core i7-10510U", RamCapacity: "8GB", RamType: "DDR4", StorageCapacity: "512GB", BatteryStatus: "Yes"}
	laptop2 := entity.Laptop{Brand: "HP", Model: "Pavilion", Processor: "AMD Ryzen 5 3500U", RamCapacity: "16GB", RamType: "DDR4", StorageCapacity: "1TB", BatteryStatus: "Yes"}

	repo.db["id1"] = laptop1
	repo.db["id2"] = laptop2

	laptops := repo.GetAll()

	assert.Len(t, laptops, 2)
	assert.Contains(t, laptops, laptop1)
	assert.Contains(t, laptops, laptop2)
}

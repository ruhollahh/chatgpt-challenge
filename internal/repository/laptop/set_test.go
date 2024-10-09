package laptoprepo

import (
	"chatgpt-challenge/internal/entity"
	"testing"
)

// Successfully adds a new laptop to the repository with a unique promptID
func TestAddLaptop(t *testing.T) {
	repo := New()
	laptop := entity.Laptop{
		Brand:           "Dell",
		Model:           "Inspiron",
		Processor:       "Intel Core i7-10510U",
		RamCapacity:     "8GB",
		RamType:         "DDR4",
		StorageCapacity: "512GB",
		BatteryStatus:   "Yes",
	}
	promptID := "unique-id-123"

	repo.Set(promptID, laptop)

	if _, exists := repo.db[promptID]; !exists {
		t.Errorf("Expected laptop to be added with promptID %s", promptID)
	}
}

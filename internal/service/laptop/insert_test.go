package laptopservice

import (
	"chatgpt-challenge/internal/entity"
	laptopparam "chatgpt-challenge/internal/param/laptop"
	"testing"
)

// Successfully inserts a laptop entity into the repository
func TestInsert(t *testing.T) {
	mockLaptopRepo := NewMockLaptopRepo(t)
	mockStructGen := NewMockStructGenerator(t)
	mockLaptopSchema := NewMockLaptopSchema(t)
	cfg := Config{SystemMessage: "Test Message"}
	service := New(cfg, mockLaptopRepo, mockStructGen, mockLaptopSchema)

	promptID := "testPromptID"
	laptop := entity.Laptop{
		Brand:           "TestBrand",
		Model:           "TestModel",
		Processor:       "TestProcessor",
		RamCapacity:     "TestRamCapacity",
		RamType:         "TestRamType",
		StorageCapacity: "TestStorageCapacity",
		BatteryStatus:   "TestBatteryStatus",
	}

	mockLaptopRepo.EXPECT().Set(promptID, laptop).Return().Once()

	service.Insert(laptopparam.InsertRequest{
		PromptID: promptID,
		Laptop:   laptop,
	})
}

package laptopservice

import (
	"chatgpt-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllReturnsCorrectNumberOfResponses(t *testing.T) {
	mockLaptopRepo := NewMockLaptopRepo(t)
	mockStructGen := NewMockStructGenerator(t)
	mockLaptopSchema := NewMockLaptopSchema(t)
	laptopService := New(Config{}, mockLaptopRepo, mockStructGen, mockLaptopSchema)

	laptops := []entity.Laptop{
		{Brand: "Dell", Model: "Inspiron", Processor: "Intel Core i7-10510U", RamCapacity: "8GB", RamType: "DDR4", StorageCapacity: "512GB", BatteryStatus: "Yes"},
		{Brand: "HP", Model: "Pavilion", Processor: "Intel Core i5-1035G1", RamCapacity: "16GB", RamType: "DDR4", StorageCapacity: "1TB", BatteryStatus: "Yes"},
	}

	mockLaptopRepo.EXPECT().GetAll().Return(laptops).Once()

	response := laptopService.GetAll()

	assert.Equal(t, len(laptops), len(response))
}

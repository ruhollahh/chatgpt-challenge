package laptopservice

import (
	laptopparam "chatgpt-challenge/internal/param/laptop"
)

func (s Service) GetAll() []laptopparam.GetAllResponse {
	laptops := s.laptopRepo.GetAll()
	response := make([]laptopparam.GetAllResponse, len(laptops))
	for i, laptop := range laptops {
		response[i] = laptopparam.GetAllResponse{
			Brand:           laptop.Brand,
			Model:           laptop.Model,
			Processor:       laptop.Processor,
			RamCapacity:     laptop.RamCapacity,
			RamType:         laptop.RamType,
			StorageCapacity: laptop.StorageCapacity,
			BatteryStatus:   laptop.BatteryStatus,
		}
	}

	return response
}

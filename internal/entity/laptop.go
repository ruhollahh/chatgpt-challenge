package entity

type Laptop struct {
	Brand           string `json:"brand" required:"true" description:"Laptop brand name, example: Dell"`
	Model           string `json:"model" required:"true" description:"Laptop model name, example: Inspiron"`
	Processor       string `json:"processor" required:"true" description:"Laptop processor must contain manufacturer and series name, example: Intel Core i7-10510U"`
	RamCapacity     string `json:"ram_capacity" required:"true" description:"Laptop RAM capacity must only contain the size of the RAM, example: 8GB"`
	RamType         string `json:"ram_type" required:"true" description:"Laptop RAM type, example: DDR4"`
	StorageCapacity string `json:"storage_capacity" required:"true" description:"Laptop storage capacity must only contain the size of the storage and NOT the type, example: 512GB"`
	BatteryStatus   string `json:"battery_status" required:"true" description:"Laptop battery status must only be No if the battery is damaged or removed otherwise Yes, example: No"`
}

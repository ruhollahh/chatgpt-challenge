package laptopservice

import (
	"chatgpt-challenge/internal/entity"
	"chatgpt-challenge/internal/param/openai"
	"encoding/json"
)

type Config struct {
	SystemMessage string
}

//go:generate mockery --name LaptopRepo
type LaptopRepo interface {
	Set(promptID string, laptop entity.Laptop)
	GetAll() []entity.Laptop
}

//go:generate mockery --name StructGenerator
type StructGenerator interface {
	StructuredOutput(request openaiparam.StructuredOutputRequest) (openaiparam.StructuredOutputResponse, error)
}

//go:generate mockery --name LaptopSchema
type LaptopSchema interface {
	json.Marshaler
	Name() string
	Description() string
	ParseContent(content string) (entity.Laptop, error)
}

type Service struct {
	cfg             Config
	laptopRepo      LaptopRepo
	structGenerator StructGenerator
	laptopSchema    LaptopSchema
}

func New(cfg Config, laptopRepo LaptopRepo, structGenerator StructGenerator, laptopSchema LaptopSchema) Service {
	return Service{
		cfg:             cfg,
		laptopRepo:      laptopRepo,
		structGenerator: structGenerator,
		laptopSchema:    laptopSchema,
	}
}

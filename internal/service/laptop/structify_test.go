package laptopservice

import (
	"chatgpt-challenge/internal/entity"
	laptopparam "chatgpt-challenge/internal/param/laptop"
	openaiparam "chatgpt-challenge/internal/param/openai"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructify(t *testing.T) {
	mockLaptopRepo := NewMockLaptopRepo(t)
	mockStructGenerator := NewMockStructGenerator(t)
	mockLaptopSchema := NewMockLaptopSchema(t)
	cfg := Config{SystemMessage: "Test System Message"}
	service := New(cfg, mockLaptopRepo, mockStructGenerator, mockLaptopSchema)
	unstructuredText := "Sample laptop description"
	schemaName := "Laptop"
	schemaDescription := "Laptop Information"

	t.Run("Valid structure", func(t *testing.T) {
		validStructuredJSON := `{"brand":"TestBrand","model":"TestModel","processor":"TestProcessor","ram_capacity":"8GB","ram_type":"DDR4","storage_capacity":"256GB","battery_status":"Yes"}`
		validParseContentResult := entity.Laptop{
			Brand:           "TestBrand",
			Model:           "TestModel",
			Processor:       "TestProcessor",
			RamCapacity:     "8GB",
			RamType:         "DDR4",
			StorageCapacity: "256GB",
			BatteryStatus:   "Yes",
		}

		mockLaptopSchema.EXPECT().Name().Return(schemaName).Once()
		mockLaptopSchema.EXPECT().Description().Return(schemaDescription).Once()
		mockLaptopSchema.EXPECT().ParseContent(validStructuredJSON).Return(validParseContentResult, nil).Once()

		mockStructGenerator.EXPECT().StructuredOutput(openaiparam.StructuredOutputRequest{
			UnstructuredText: unstructuredText,
			JSONSchema: openaiparam.JSONSchema{
				Name:        schemaName,
				Description: schemaDescription,
				Schema:      mockLaptopSchema,
			},
			SystemMessage: cfg.SystemMessage,
		}).Return(openaiparam.StructuredOutputResponse{StructuredJSON: validStructuredJSON}, nil).Once()

		resp, err := service.Structify(laptopparam.StructifyRequest{
			Content: unstructuredText,
		})

		assert.NoError(t, err)
		assert.Equal(t, "TestBrand", resp.Laptop.Brand)
		assert.Equal(t, "TestModel", resp.Laptop.Model)
	})

	t.Run("Invalid structure", func(t *testing.T) {
		invalidStructuredJSON := `{"invalid_key":"invalid","model":"TestModel","processor":"TestProcessor","ram_capacity":"8GB","ram_type":"DDR4","storage_capacity":"256GB","battery_status":"Yes"}`
		parseContentResult := entity.Laptop{}

		mockLaptopSchema.EXPECT().Name().Return(schemaName).Once()
		mockLaptopSchema.EXPECT().Description().Return(schemaDescription).Once()
		mockLaptopSchema.EXPECT().ParseContent(invalidStructuredJSON).Return(parseContentResult, errors.New("invalid json")).Once()

		mockStructGenerator.EXPECT().StructuredOutput(openaiparam.StructuredOutputRequest{
			UnstructuredText: unstructuredText,
			JSONSchema: openaiparam.JSONSchema{
				Name:        schemaName,
				Description: schemaDescription,
				Schema:      mockLaptopSchema,
			},
			SystemMessage: cfg.SystemMessage,
		}).Return(openaiparam.StructuredOutputResponse{StructuredJSON: invalidStructuredJSON}, nil).Once()

		_, err := service.Structify(laptopparam.StructifyRequest{
			Content: unstructuredText,
		})

		assert.Error(t, err)
	})

	t.Run("Invalid structure content", func(t *testing.T) {
		invalidStructuredJSONContent := `{"brand":"","model":"TestModel","processor":"TestProcessor","ram_capacity":"8GB","ram_type":"DDR4","storage_capacity":"256GB","battery_status":"Yes"}`
		parseContentResult := entity.Laptop{}

		mockLaptopSchema.EXPECT().Name().Return(schemaName).Once()
		mockLaptopSchema.EXPECT().Description().Return(schemaDescription).Once()
		mockLaptopSchema.EXPECT().ParseContent(invalidStructuredJSONContent).Return(parseContentResult, errors.New("invalid json")).Once()

		mockStructGenerator.EXPECT().StructuredOutput(openaiparam.StructuredOutputRequest{
			UnstructuredText: unstructuredText,
			JSONSchema: openaiparam.JSONSchema{
				Name:        schemaName,
				Description: schemaDescription,
				Schema:      mockLaptopSchema,
			},
			SystemMessage: cfg.SystemMessage,
		}).Return(openaiparam.StructuredOutputResponse{StructuredJSON: invalidStructuredJSONContent}, nil).Once()

		_, err := service.Structify(laptopparam.StructifyRequest{
			Content: unstructuredText,
		})

		assert.Error(t, err)
	})
}

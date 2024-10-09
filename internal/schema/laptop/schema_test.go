package laptopschema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchema_ParseContent(t *testing.T) {
	t.Run("With valid JSON", func(t *testing.T) {

		schema := New()
		jsonContent := `{
			"brand": "Dell",
			"model": "Inspiron",
			"processor": "Intel Core i7-10510U",
			"ram_capacity": "8GB",
			"ram_type": "DDR4",
			"storage_capacity": "512GB",
			"battery_status": "Yes"
    	}`

		laptop, err := schema.ParseContent(jsonContent)
		assert.NoError(t, err)

		assert.Equal(t, "Dell", laptop.Brand)
		assert.Equal(t, "Inspiron", laptop.Model)
		assert.Equal(t, "Intel Core i7-10510U", laptop.Processor)
		assert.Equal(t, "8GB", laptop.RamCapacity)
		assert.Equal(t, "DDR4", laptop.RamType)
		assert.Equal(t, "512GB", laptop.StorageCapacity)
		assert.Equal(t, "Yes", laptop.BatteryStatus)
	})

	t.Run("With invalid JSON", func(t *testing.T) {
		schema := New()
		jsonContent := `{"invalid_key":"invalid","model":"TestModel","processor":"TestProcessor","ram_capacity":"8GB","ram_type":"DDR4","storage_capacity":"256GB","battery_status":"Yes"}`

		_, err := schema.ParseContent(jsonContent)
		assert.Error(t, err)
	})

	t.Run("With empty JSON", func(t *testing.T) {
		schema := New()
		jsonContent := ""

		_, err := schema.ParseContent(jsonContent)
		assert.Error(t, err)
	})
}

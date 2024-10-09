package promptrepo

import (
	"chatgpt-challenge/internal/entity"
	"testing"
)

func TestSetNX(t *testing.T) {
	t.Run("Adds new prompt", func(t *testing.T) {

		repo := New()
		prompt := entity.Prompt{
			ID:      "1",
			Content: "Test content",
			Status:  entity.PromptStatusPending,
		}

		success := repo.SetNX(prompt)

		if !success {
			t.Errorf("Expected true, got false")
		}

		if _, exists := repo.db[prompt.ID]; !exists {
			t.Errorf("Expected prompt to be added to the repository")
		}
	})

	t.Run("Returns false for existing ID", func(t *testing.T) {
		repo := New()
		prompt := entity.Prompt{
			ID:      "1",
			Content: "Test content",
			Status:  entity.PromptStatusPending,
		}

		repo.SetNX(prompt)
		success := repo.SetNX(prompt)

		if success {
			t.Errorf("Expected false, got true")
		}
	})
}

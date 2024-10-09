package promptrepo

import (
	"chatgpt-challenge/internal/entity"
	"chatgpt-challenge/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateStatus(t *testing.T) {
	t.Run("Status changes successfully", func(t *testing.T) {
		repo := New()
		promptID := "123"
		initialPrompt := entity.Prompt{
			ID:      promptID,
			Content: "Test content",
			Status:  entity.PromptStatusPending,
		}
		repo.db[promptID] = initialPrompt

		err := repo.UpdateStatus(promptID, entity.PromptStatusProcessed)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		updatedPrompt, exists := repo.db[promptID]

		assert.True(t, exists)
		assert.Equal(t, updatedPrompt.Status, entity.PromptStatusProcessed)
	})

	t.Run("Error prompt not found if prompt doesn't exist", func(t *testing.T) {
		repo := New()
		nonExistentID := "999"

		err := repo.UpdateStatus(nonExistentID, entity.PromptStatusProcessed)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrNotFound)
	})
}

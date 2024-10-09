package promptrepo

import (
	"chatgpt-challenge/internal/entity"
	"chatgpt-challenge/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateErrorMessage(t *testing.T) {
	t.Run("Update is successful", func(t *testing.T) {
		repo := New()
		promptID := "123"
		initialPrompt := entity.Prompt{
			ID:           promptID,
			Content:      "Test content",
			Status:       entity.PromptStatusPending,
			ErrorMessage: "",
		}
		repo.db[promptID] = initialPrompt

		errMessage := "something went wrong"
		err := repo.UpdateErrorMessage(promptID, errMessage)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		updatedPrompt, exists := repo.db[promptID]
		if !exists {
			t.Fatalf("expected prompt to exist in the repository")
		}

		assert.Equal(t, updatedPrompt.ErrorMessage, errMessage)
	})

	t.Run("Error prompt not found if prompt doesn't exist", func(t *testing.T) {
		repo := New()
		nonExistentID := "999"

		err := repo.UpdateStatus(nonExistentID, entity.PromptStatusProcessed)
		assert.NotNil(t, err)
		assert.Equal(t, err, repository.ErrNotFound)
	})
}

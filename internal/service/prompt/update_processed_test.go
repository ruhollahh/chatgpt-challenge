package promptservice

import (
	"chatgpt-challenge/internal/entity"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateProcessed(t *testing.T) {
	mockPromptRepo := NewMockPromptRepo(t)
	service := New(mockPromptRepo)

	t.Run("Successful processed update", func(t *testing.T) {
		req := promptparam.UpdateProcessedRequest{
			ID: "valid-id",
		}

		mockPromptRepo.EXPECT().UpdateStatus(req.ID, entity.PromptStatusProcessed).Return(nil).Once()

		err := service.UpdateProcessed(req)

		assert.NoError(t, err)
	})

	t.Run("Error processed update", func(t *testing.T) {
		req := promptparam.UpdateProcessedRequest{
			ID: "valid-id",
		}

		mockPromptRepo.EXPECT().UpdateStatus(req.ID, entity.PromptStatusProcessed).Return(errors.New("some error")).Once()

		err := service.UpdateProcessed(req)

		assert.Error(t, err)
	})
}

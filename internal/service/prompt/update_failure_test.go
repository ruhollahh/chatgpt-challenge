package promptservice

import (
	"chatgpt-challenge/internal/entity"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateFailure(t *testing.T) {
	mockPromptRepo := NewMockPromptRepo(t)
	service := New(mockPromptRepo)

	t.Run("Successful failure update", func(t *testing.T) {
		req := promptparam.UpdateFailureRequest{
			ID:           "valid-id",
			ErrorMessage: "Some error occurred",
		}

		mockPromptRepo.EXPECT().UpdateStatus(req.ID, entity.PromptStatusFailed).Return(nil).Once()
		mockPromptRepo.EXPECT().UpdateErrorMessage(req.ID, req.ErrorMessage).Return(nil).Once()

		err := service.UpdateFailure(req)

		assert.NoError(t, err)
	})

	t.Run("Error failure update", func(t *testing.T) {
		req := promptparam.UpdateFailureRequest{
			ID:           "invalid-id",
			ErrorMessage: "Some error occurred",
		}

		mockPromptRepo.EXPECT().UpdateStatus(req.ID, entity.PromptStatusFailed).Return(errors.New("some error")).Once()

		err := service.UpdateFailure(req)

		assert.Error(t, err)
	})
}

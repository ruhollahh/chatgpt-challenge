package promptservice

import (
	"chatgpt-challenge/internal/entity"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertIfNotExist(t *testing.T) {
	mockPromptRepo := NewMockPromptRepo(t)
	service := New(mockPromptRepo)

	t.Run("Successfully insert", func(t *testing.T) {
		req := promptparam.InsertIfNotExistRequest{
			ID:      "123",
			Content: "Test Content",
		}

		mockPromptRepo.EXPECT().SetNX(entity.Prompt{
			ID:           req.ID,
			Content:      req.Content,
			Status:       entity.PromptStatusPending,
			ErrorMessage: "",
		}).Return(true).Once()

		response := service.InsertIfNotExist(req)

		assert.True(t, response.Inserted)
	})

	t.Run("Fail to insert duplicate", func(t *testing.T) {
		req := promptparam.InsertIfNotExistRequest{
			ID:      "123",
			Content: "Test Content",
		}

		mockPromptRepo.EXPECT().SetNX(entity.Prompt{
			ID:           req.ID,
			Content:      req.Content,
			Status:       entity.PromptStatusPending,
			ErrorMessage: "",
		}).Return(false).Once()

		response := service.InsertIfNotExist(req)

		assert.False(t, response.Inserted)
	})
}

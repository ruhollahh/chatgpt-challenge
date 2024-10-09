package promptservice

import (
	"chatgpt-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllReturnsCorrectNumberOfResponses(t *testing.T) {
	mockPromptRepo := NewMockPromptRepo(t)
	promptService := New(mockPromptRepo)

	prompts := []entity.Prompt{
		{
			ID:           "1",
			Content:      "Test Content",
			Status:       entity.PromptStatusPending,
			ErrorMessage: "",
		},
		{
			ID:           "2",
			Content:      "Test Content",
			Status:       entity.PromptStatusPending,
			ErrorMessage: "",
		},
	}

	mockPromptRepo.EXPECT().GetAll().Return(prompts).Once()

	response := promptService.GetAll()

	assert.Equal(t, len(prompts), len(response))
}

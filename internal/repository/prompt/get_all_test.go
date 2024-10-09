package promptrepo

import (
	"chatgpt-challenge/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Retrieve all prompts when the repository is populated
func TestGetAll(t *testing.T) {
	repo := New()
	prompt1 := entity.Prompt{ID: "1", Content: "Prompt 1", Status: entity.PromptStatusPending}
	prompt2 := entity.Prompt{ID: "2", Content: "Prompt 2", Status: entity.PromptStatusPending}

	repo.db["1"] = prompt1
	repo.db["2"] = prompt2

	prompts := repo.GetAll()

	assert.Len(t, prompts, 2)
	assert.Contains(t, prompts, prompt1)
	assert.Contains(t, prompts, prompt2)
}

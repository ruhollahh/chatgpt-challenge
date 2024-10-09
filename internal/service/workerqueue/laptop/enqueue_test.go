package laptopworkerqueue

import (
	promptparam "chatgpt-challenge/internal/param/prompt"
	"crypto/sha256"
	"encoding/base32"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestEnqueue(t *testing.T) {
	mockPromptService := NewMockPromptService(t)
	mockLaptopService := NewMockLaptopService(t)
	cfg := Config{BufferSize: 5, Workers: 3}
	workerQueue := New(cfg, mockPromptService, mockLaptopService)

	t.Run("Task is enqueued when SetNX returns true", func(t *testing.T) {
		task := Task{PromptContent: "Test content"}
		hashedPrompt := sha256.Sum256([]byte(task.PromptContent))
		promptID := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashedPrompt[:])
		mockPromptService.EXPECT().InsertIfNotExist(promptparam.InsertIfNotExistRequest{
			ID:      promptID,
			Content: task.PromptContent,
		}).Return(promptparam.InsertIfNotExistResponse{
			Inserted: true,
		}).Once()

		workerQueue.Enqueue(task)

		select {
		case enqueuedTask := <-workerQueue.tasks:
			assert.Equal(t, task.PromptContent, enqueuedTask.PromptContent)
		default:
			t.Error("Task was not enqueued")
		}
	})

	t.Run("Task is not enqueued when SetNX returns false", func(t *testing.T) {
		mockPromptService.EXPECT().InsertIfNotExist(mock.Anything).Return(promptparam.InsertIfNotExistResponse{
			Inserted: false,
		}).Once()

		workerQueue.Enqueue(Task{PromptContent: "Test content"})
		select {
		case enqueuedTask := <-workerQueue.tasks:
			assert.Fail(t, "Expected channel to be empty, but received value", enqueuedTask)
		default:
			assert.True(t, true, "Channel is empty")
		}
	})
}

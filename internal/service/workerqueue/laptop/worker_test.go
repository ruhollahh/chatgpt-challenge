package laptopworkerqueue

import (
	"chatgpt-challenge/internal/entity"
	laptopparam "chatgpt-challenge/internal/param/laptop"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	cfg := Config{BufferSize: 10, Workers: 1}
	mockPromptService := NewMockPromptService(t)
	mockLaptopService := NewMockLaptopService(t)

	t.Run("Processes tasks without errors", func(t *testing.T) {
		workerQueue := New(cfg, mockPromptService, mockLaptopService)
		task := Task{
			promptID:      "123",
			PromptContent: "test content",
		}
		workerQueue.wg.Add(1)
		workerQueue.tasks <- task

		mockLaptopService.EXPECT().Structify(laptopparam.StructifyRequest{Content: "test content"}).
			Return(laptopparam.StructifyResponse{Laptop: entity.Laptop{}}, nil).Once()
		mockPromptService.EXPECT().UpdateProcessed(promptparam.UpdateProcessedRequest{ID: "123"}).Return(nil).Once()
		mockLaptopService.EXPECT().Insert(laptopparam.InsertRequest{PromptID: "123", Laptop: entity.Laptop{}}).Return().Once()

		go workerQueue.worker()
		workerQueue.GracefullyStop()
	})

	t.Run("Processes tasks with errors", func(t *testing.T) {
		workerQueue := New(cfg, mockPromptService, mockLaptopService)
		task := Task{
			promptID:      "123",
			PromptContent: "test content",
		}
		workerQueue.wg.Add(1)
		workerQueue.tasks <- task

		structifyErr := fmt.Errorf("structify error")
		mockLaptopService.EXPECT().Structify(laptopparam.StructifyRequest{Content: "test content"}).
			Return(laptopparam.StructifyResponse{}, structifyErr).Once()
		mockPromptService.EXPECT().UpdateFailure(promptparam.UpdateFailureRequest{
			ID:           "123",
			ErrorMessage: fmt.Sprintf("error structifying prompt content: %s\n", structifyErr.Error()),
		}).Return(nil).Once()

		go workerQueue.worker()
		workerQueue.GracefullyStop()
	})
}

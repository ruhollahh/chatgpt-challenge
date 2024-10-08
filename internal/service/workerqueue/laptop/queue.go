package laptopworkerqueue

import (
	"chatgpt-challenge/internal/entity"
	laptopstructifyparam "chatgpt-challenge/internal/param/structify/laptop"
	"fmt"
	"sync"
)

type Config struct {
	BufferSize int
	Workers    int
}

type PromptRepository interface {
	SetNX(prompt entity.Prompt) bool
	UpdateStatus(id string, status entity.PromptStatus) error
}

type LaptopRepository interface {
	Set(promptID string, laptop entity.Laptop)
}

type LaptopStructifyService interface {
	Structify(req laptopstructifyparam.StructifyRequest) (laptopstructifyparam.StructifyResponse, error)
}

type WorkerQueue struct {
	cfg                    Config
	tasks                  chan Task
	laptopRepo             LaptopRepository
	promptRepo             PromptRepository
	laptopStructifyService LaptopStructifyService
	wg                     *sync.WaitGroup
}

type Task struct {
	PromptID      string
	PromptContent string
}

func New(cfg Config, repo LaptopRepository, promptRepo PromptRepository, laptopStructifySvc LaptopStructifyService) WorkerQueue {
	return WorkerQueue{
		cfg:                    cfg,
		tasks:                  make(chan Task, cfg.BufferSize),
		laptopRepo:             repo,
		promptRepo:             promptRepo,
		laptopStructifyService: laptopStructifySvc,
		wg:                     &sync.WaitGroup{},
	}
}

func (q *WorkerQueue) Start() {
	for i := 0; i < q.cfg.Workers; i++ {
		go q.worker()
	}
}

func (q *WorkerQueue) GracefullyStop() {
	q.wg.Wait()
	close(q.tasks)
}

func (q *WorkerQueue) worker() {
	for task := range q.tasks {
		res, err := q.laptopStructifyService.Structify(laptopstructifyparam.StructifyRequest{
			Content: task.PromptContent,
		})
		if err != nil {
			fmt.Printf("error structifying prompt content: %s\n", err.Error())
			continue
		}

		err = q.promptRepo.UpdateStatus(task.PromptID, entity.PromptStatusProcessed)
		if err != nil {
			fmt.Printf("error updating prompt status to processed: %s\n", err.Error())
			continue
		}

		q.laptopRepo.Set(task.PromptID, res.Laptop)

		q.wg.Done()
	}
}

func (q *WorkerQueue) Enqueue(task Task) {
	if hasSet := q.promptRepo.SetNX(entity.Prompt{
		ID:      task.PromptID,
		Content: task.PromptContent,
		Status:  entity.PromptStatusPending,
	}); hasSet {
		q.wg.Add(1)
		q.tasks <- task
	}
}

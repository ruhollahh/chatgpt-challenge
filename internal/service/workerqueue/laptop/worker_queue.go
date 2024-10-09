package laptopworkerqueue

import (
	laptopparam "chatgpt-challenge/internal/param/laptop"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"sync"
)

// Config contains the configuration settings for the WorkerQueue system, including the buffer size
// for the task channel and the number of worker goroutines to spawn.
type Config struct {
	BufferSize int
	Workers    int
}

//go:generate mockery --name PromptService
type PromptService interface {
	InsertIfNotExist(req promptparam.InsertIfNotExistRequest) promptparam.InsertIfNotExistResponse
	UpdateFailure(req promptparam.UpdateFailureRequest) error
	UpdateProcessed(req promptparam.UpdateProcessedRequest) error
}

//go:generate mockery --name LaptopService
type LaptopService interface {
	Insert(request laptopparam.InsertRequest)
	Structify(laptopparam.StructifyRequest) (laptopparam.StructifyResponse, error)
}

// WorkerQueue manages the task queue and workers for processing prompts and structuring them into laptop entities.
type WorkerQueue struct {
	cfg           Config
	tasks         chan Task
	promptService PromptService
	laptopService LaptopService
	wg            *sync.WaitGroup
}

// Task represents a unit of work for the WorkerQueue.
type Task struct {
	promptID      string
	PromptContent string
}

func New(cfg Config, promptService PromptService, laptopSvc LaptopService) WorkerQueue {
	return WorkerQueue{
		cfg:           cfg,
		tasks:         make(chan Task, cfg.BufferSize),
		promptService: promptService,
		laptopService: laptopSvc,
		wg:            &sync.WaitGroup{},
	}
}

// Start launches the worker goroutines based on the configured number of workers.
func (q *WorkerQueue) Start() {
	for i := 0; i < q.cfg.Workers; i++ {
		go q.worker()
	}
}

// GracefullyStop waits for all currently queued tasks to complete and then closes the task channel.
func (q *WorkerQueue) GracefullyStop() {
	q.wg.Wait()
	close(q.tasks)
}

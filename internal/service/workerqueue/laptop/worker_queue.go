package laptopworkerqueue

import (
	laptopparam "chatgpt-challenge/internal/param/laptop"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"sync"
)

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

type WorkerQueue struct {
	cfg           Config
	tasks         chan Task
	promptService PromptService
	laptopService LaptopService
	wg            *sync.WaitGroup
}

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

func (q *WorkerQueue) Start() {
	for i := 0; i < q.cfg.Workers; i++ {
		go q.worker()
	}
}

func (q *WorkerQueue) GracefullyStop() {
	q.wg.Wait()
	close(q.tasks)
}

package laptopworkerqueue

import (
	laptopparam "chatgpt-challenge/internal/param/laptop"
	promptparam "chatgpt-challenge/internal/param/prompt"
	"fmt"
	"log"
)

func (q *WorkerQueue) worker() {
	for task := range q.tasks {
		res, err := q.laptopService.Structify(laptopparam.StructifyRequest{
			Content: task.PromptContent,
		})
		if err != nil {
			log.Printf("error structifying prompt content: %s\n", err.Error())
			err = q.promptService.UpdateFailure(promptparam.UpdateFailureRequest{
				ID:           task.promptID,
				ErrorMessage: fmt.Sprintf("error structifying prompt content: %s\n", err.Error()),
			})
			if err != nil {
				log.Printf("error updating prompt status to failed: %s\n", err.Error())
			}
			q.wg.Done()
			continue
		}

		err = q.promptService.UpdateProcessed(promptparam.UpdateProcessedRequest{
			ID: task.promptID,
		})
		if err != nil {
			log.Printf("error updating prompt status to processed: %s\n", err.Error())
			err = q.promptService.UpdateFailure(promptparam.UpdateFailureRequest{
				ID:           task.promptID,
				ErrorMessage: fmt.Sprintf("error updating prompt to processed: %s", err.Error()),
			})
			if err != nil {
				log.Printf("error updating prompt status to failed: %s\n", err.Error())
			}
			q.wg.Done()
			continue
		}

		q.laptopService.Insert(laptopparam.InsertRequest{
			PromptID: task.promptID,
			Laptop:   res.Laptop,
		})

		q.wg.Done()
	}
}

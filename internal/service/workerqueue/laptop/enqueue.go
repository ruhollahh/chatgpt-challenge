package laptopworkerqueue

import (
	promptparam "chatgpt-challenge/internal/param/prompt"
	"crypto/sha256"
	"encoding/base32"
)

func (q *WorkerQueue) Enqueue(task Task) {
	hashedPrompt := sha256.Sum256([]byte(task.PromptContent))
	task.promptID = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashedPrompt[:])

	if res := q.promptService.InsertIfNotExist(promptparam.InsertIfNotExistRequest{
		ID:      task.promptID,
		Content: task.PromptContent,
	}); res.Inserted {
		q.wg.Add(1)
		q.tasks <- task
	}
}

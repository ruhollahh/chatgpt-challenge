package main

import (
	openaiclient "chatgpt-challenge/internal/client/openai"
	laptoprepo "chatgpt-challenge/internal/repository/laptop"
	promptrepo "chatgpt-challenge/internal/repository/prompt"
	laptopstructify "chatgpt-challenge/internal/service/structify/laptop"
	laptopworkerqueue "chatgpt-challenge/internal/service/workerqueue/laptop"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

const laptopSchemaSystemMessage = "You are an expert in structured data extraction. You will be provided with unstructured text describing a laptop and you should convert it into the specified structure. Fill in any required missing information using accurate data and disregard any extra details."

func main() {
	authToken := os.Getenv("CHATGPT_CHALLENGE__AUTH_TOKEN")

	openAiClient := openaiclient.New(openaiclient.Config{
		AuthToken:            authToken,
		Model:                openai.GPT4oMini,
		MaxCompletionsTokens: 60,
	})
	laptopStructifySvc := laptopstructify.New(laptopstructify.Config{
		SystemMessage: laptopSchemaSystemMessage,
	}, openAiClient)

	promptRepo := promptrepo.New()
	laptopRepo := laptoprepo.New()

	startLaptopWorkerQueue(laptopRepo, promptRepo, laptopStructifySvc)

	fmt.Printf("Prompts: %+v\n", promptRepo.GetAll())
	fmt.Printf("Laptops: %+v\n", laptopRepo.GetAll())
}

func startLaptopWorkerQueue(laptopRepo laptopworkerqueue.LaptopRepository, promptRepo laptopworkerqueue.PromptRepository, laptopStructifySvc laptopworkerqueue.LaptopStructifyService) {
	prompts := []string{
		"Laptop: Dell Inspiron; Processor i7-10510U ; RAM 16GB; 512GB SSD Missing battery",
		"MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed",
		"ThinkPad, i5 CPU, 8GB memory, storage: 1TB HDD",
		"Asus ROG, Processor: AMD Ryzen 7; RAM 16 GB; 1TB SSD; Damaged battrey",
		"Dell Inspiron; Processor: i5-1135G7; RAM 8GB; Storage: 256.123548 SSD; Missing charger",
		"Laptop: Dell Inspiron; Processor i7-10510U ; RAM 16GB; 512GB SSD Missing battery",
		"MacBook Pro with M1 chip, 8GB RAM, 256 GB SSD storage Battery removed",
		"ThinkPad, i5 CPU, 8GB memory, storage: 1TB HDD",
		"Asus ROG, Processor: AMD Ryzen 7; RAM 16 GB; 1TB SSD; Damaged battrey",
		"Dell Inspiron; Processor: i5-1135G7; RAM 8GB; Storage: 256.123548 SSD; Missing charger",
	}

	laptopWorkerQueue := laptopworkerqueue.New(
		laptopworkerqueue.Config{
			BufferSize: 5,
			Workers:    3,
		},
		laptopRepo,
		promptRepo,
		laptopStructifySvc,
	)

	laptopWorkerQueue.Start()
	defer laptopWorkerQueue.GracefullyStop()

	for _, prompt := range prompts {
		hashedPrompt := sha256.Sum256([]byte(prompt))
		promptID := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hashedPrompt[:])
		laptopWorkerQueue.Enqueue(laptopworkerqueue.Task{
			PromptID:      promptID,
			PromptContent: prompt,
		})
	}
}

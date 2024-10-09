package main

import (
	"chatgpt-challenge/config"
	httpserver "chatgpt-challenge/delivery/http_server"
	_ "chatgpt-challenge/docs"
	openaiclient "chatgpt-challenge/internal/client/openai"
	laptoprepo "chatgpt-challenge/internal/repository/laptop"
	promptrepo "chatgpt-challenge/internal/repository/prompt"
	"chatgpt-challenge/internal/schema/laptop"
	"chatgpt-challenge/internal/service/laptop"
	promptservice "chatgpt-challenge/internal/service/prompt"
	laptopworkerqueue "chatgpt-challenge/internal/service/workerqueue/laptop"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

const laptopSchemaSystemMessage = "You are an expert in structured data extraction. You will be provided with unstructured text describing a laptop and you should convert it into the specified structure. Fill in any required missing information using accurate data and disregard any extra details."

// @title ChatGPT Challenge
// @version 1.0
// @description This is an easy way to retrieve the generated structured data.
// @termsOfService https://example.com/terms

// @contact.name API Support
// @contact.url https://www.example.com/support
// @contact.email ruhollahh01@gmail.com

// @BasePath /
func main() {
	authToken := os.Getenv("CHATGPT_CHALLENGE__AUTH_TOKEN")

	cfg := config.Config{
		HTTPServer: httpserver.Config{
			Port: 1414,
		},
		OpenAIClient: openaiclient.Config{
			AuthToken:            authToken,
			Model:                openai.GPT4oMini,
			MaxCompletionsTokens: 60,
		},
		LaptopService: laptopservice.Config{
			SystemMessage: laptopSchemaSystemMessage,
		},
		LaptopWorkerQueue: laptopworkerqueue.Config{
			BufferSize: 5,
			Workers:    3,
		},
	}

	openAIClient := openaiclient.New(cfg.OpenAIClient)

	promptRepo := promptrepo.New()
	laptopRepo := laptoprepo.New()

	promptSvc := promptservice.New(promptRepo)
	laptopSvc := laptopservice.New(
		cfg.LaptopService,
		laptopRepo,
		openAIClient,
		laptopschema.New(),
	)

	laptopWorkerQueue := laptopworkerqueue.New(
		cfg.LaptopWorkerQueue,
		promptSvc,
		laptopSvc,
	)

	go startLaptopWorkerQueue(laptopWorkerQueue)

	httpServerConfig := httpserver.Config{Port: 1414}
	server := httpserver.New(httpServerConfig, promptSvc, laptopSvc)
	server.RegisterRoutes()

	log.Println("starting server", "port", httpServerConfig.Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// startLaptopWorkerQueue starts the laptop worker queue, enqueues a list of unstructured
// laptop descriptions for structuring, and gracefully stops the worker queue after
// processing all the tasks.
//
// Parameters:
// - workerQueue: The worker queue responsible for structuring laptop descriptions.
func startLaptopWorkerQueue(workerQueue laptopworkerqueue.WorkerQueue) {
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

	workerQueue.Start()
	defer workerQueue.GracefullyStop()

	for _, prompt := range prompts {
		workerQueue.Enqueue(laptopworkerqueue.Task{
			PromptContent: prompt,
		})
	}
}

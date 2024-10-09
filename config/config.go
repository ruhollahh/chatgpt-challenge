package config

import (
	httpserver "chatgpt-challenge/delivery/http_server"
	openaiclient "chatgpt-challenge/internal/client/openai"
	laptopservice "chatgpt-challenge/internal/service/laptop"
	laptopworkerqueue "chatgpt-challenge/internal/service/workerqueue/laptop"
)

type Config struct {
	HTTPServer        httpserver.Config
	OpenAIClient      openaiclient.Config
	LaptopService     laptopservice.Config
	LaptopWorkerQueue laptopworkerqueue.Config
}

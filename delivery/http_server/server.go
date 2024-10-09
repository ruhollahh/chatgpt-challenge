package httpserver

import (
	"chatgpt-challenge/delivery/http_server/http_io"
	laptophandler "chatgpt-challenge/delivery/http_server/laptop_handler"
	"chatgpt-challenge/delivery/http_server/middleware"
	prompthandler "chatgpt-challenge/delivery/http_server/prompt_handler"
	laptopservice "chatgpt-challenge/internal/service/laptop"
	promptservice "chatgpt-challenge/internal/service/prompt"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Port int
}

type Server struct {
	*http.Server
	cfg       Config
	router    *http.ServeMux
	promptSvc promptservice.Service
	laptopSvc laptopservice.Service
}

func New(cfg Config, promptSvc promptservice.Service, laptopSvc laptopservice.Service) Server {
	router := http.NewServeMux()
	ms := middleware.CreateStack(middleware.Recovering, middleware.Logging)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      ms(router),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     log.Default(),
	}

	return Server{
		Server:    srv,
		cfg:       cfg,
		router:    router,
		promptSvc: promptSvc,
		laptopSvc: laptopSvc,
	}
}

func (s Server) RegisterRoutes() {
	s.router.HandleFunc("GET /health-check", healthCheck)

	s.router.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	laptopHandler := laptophandler.New(s.laptopSvc)
	laptopHandler.RegisterRoutes(s.router)

	promptHandler := prompthandler.New(s.promptSvc)
	promptHandler.RegisterRoutes(s.router)
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	http_io.WriteJSON(w, http.StatusOK, http_io.Envelope{Data: "all good"}, nil)
}

package prompthandler

import "net/http"

func (h Handler) RegisterRoutes(router *http.ServeMux) {
	user := http.NewServeMux()
	user.HandleFunc("GET /", h.GetAll)

	router.Handle("/prompts/", http.StripPrefix("/prompts", user))
}

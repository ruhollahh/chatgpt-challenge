package laptophandler

import "net/http"

func (h Handler) RegisterRoutes(router *http.ServeMux) {
	user := http.NewServeMux()
	user.HandleFunc("GET /", h.GetAll)

	router.Handle("/laptops/", http.StripPrefix("/laptops", user))
}

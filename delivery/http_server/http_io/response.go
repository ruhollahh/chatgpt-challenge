package http_io

import (
	"encoding/json"
	"log"
	"net/http"
)

type Envelope struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, data Envelope, headers http.Header) {
	js, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to marshal data: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(js)
	if err != nil {
		log.Printf("failed to write response: %v\n", err)
	}
}

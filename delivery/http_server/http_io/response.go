package http_io

import (
	"encoding/json"
	"log"
	"net/http"
)

// Envelope is a generic wrapper for API responses. It holds the response data
// and potential errors. Both fields are optional and will be omitted from the
// JSON output if they are empty.
type Envelope struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
}

// WriteJSON serializes the given data as JSON and writes it to the provided
// http.ResponseWriter with the specified HTTP status code. Optionally, custom
// headers can be added to the response.
//
// If the data cannot be marshaled into JSON, the function logs the error and
// sends a 500 Internal Server Error response to the client. It also ensures
// the JSON output ends with a newline character for better readability.
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

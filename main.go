package main

import (
	"encoding/json"
	"net/http"

	"github.com/ibzsy/cardboard/repl"
)

// RequestBody represents the structure of the input JSON
type RequestBody struct {
	Input string `json:"input"`
}

// ResponseBody represents the structure of the output JSON
type ResponseBody struct {
	Output string `json:"output"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Echo back the received string
	respBody := ResponseBody{
		Output: repl.StartREPL(Input),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

func main() {
	http.HandleFunc("/repl", handleRequest)
	port := ":8080"
	println("Server Listening On Port ->", port)
	http.ListenAndServe(port, nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string   `json:"message"`
	Page    string   `json:"page"`
	Input   string   `json:"input"`
	Words   []string `json:"words"`
}

func words(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Input := r.URL.Query().Get("input")
	msg := Response{Message: "Hello, World!", Page: "words", Input: Input, Words: []string{Input}}
	json.NewEncoder(w).Encode(msg)
}

func main() {
	http.HandleFunc("/", words)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{Message: "OK"})
	})

	fmt.Println("Server is running on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

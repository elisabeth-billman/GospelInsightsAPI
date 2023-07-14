package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Enable CORS middleware
	r.Use(corsMiddleware)

	// Define your API routes and handlers
	r.HandleFunc("/ask", handleAsk)

	port := "8000"
	fmt.Print("server is running..")
	http.ListenAndServe(":"+port, r)
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleAsk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	url := "https://api.openai.com/v1/chat/completions"
	apiKey := "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"

	// Read the request payload
	var payload map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error decoding request payload: %v", err)
		return
	}

	// Prepare the request payload
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshaling request payload: %v", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating HTTP request: %v", err)
		return
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error sending HTTP request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading response body: %v", err)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	// Write the response
	_, err = w.Write(respBody)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

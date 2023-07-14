package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestHandleAsk(t *testing.T) {
	// Create a sample request payload
	payload := map[string]interface{}{
		"question": "What is the meaning of life?",
	}

	// Convert the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Error marshaling request payload: %v", err)
	}

	// Create a mock HTTP request with the payload
	req := httptest.NewRequest(http.MethodPost, "/ask", bytes.NewBuffer(payloadBytes))

	// Create a mock HTTP response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function
	handleAsk(recorder, req)

	// Get the response from the recorder
	resp := recorder.Result()

	// Assert the response status code using the assert package
	assert.Equal(t, 400, resp.StatusCode, "Unexpected status code")
	assert.NotEqual(t, http.StatusOK, resp.StatusCode, "Unexpected status code")

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	// Perform assertions on the response body
	expectedResponseBody := []byte("{\n    \"error\": {\n        \"message\": \"you must provide a model parameter\",\n        \"type\": \"invalid_request_error\",\n        \"param\": null,\n        \"code\": null\n    }\n}\n")
	assert.Equal(t, expectedResponseBody, respBody, "Unexpected response body")
	assert.NotEqual(t, "bad response body", respBody, "Unexpected response body")
	assert.NotContains(t, string(respBody), "unexpected substring", "Unexpected substring found in response body")





	// Close the response body
	resp.Body.Close()
}



func TestCorsMiddleware(t *testing.T) {
	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a response recorder to record the server's response
	rr := httptest.NewRecorder()

	// Create a test handler that will be wrapped by the CORS middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check that the CORS headers are set by the middleware
		origin := w.Header().Get("Access-Control-Allow-Origin")
		if origin != "*" {
			t.Errorf("Expected Access-Control-Allow-Origin header to be *, but got %s", origin)
		}
		methods := w.Header().Get("Access-Control-Allow-Methods")
		if methods != "GET, POST, OPTIONS" {
			t.Errorf("Expected Access-Control-Allow-Methods header to be GET, POST, OPTIONS, but got %s", methods)
		}
		headers := w.Header().Get("Access-Control-Allow-Headers")
		if headers != "Content-Type" {
			t.Errorf("Expected Access-Control-Allow-Headers header to be Content-Type, but got %s", headers)
		}
	})

	// Wrap the test handler with the CORS middleware
	corsMiddleware(handler).ServeHTTP(rr, req)
}


package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read port from PORT env var (default: 8080)
	port := getEnv("PORT", "8080")

	// Register handlers
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/proxy", proxyHandler)

	// Start server
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

// helloHandler responds with a simple greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello"))
}

// proxyHandler calls an external service defined by SERVICE_URL and proxies its response.
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := os.Getenv("SERVICE_URL")
	if targetURL == "" {
		http.Error(w, "SERVICE_URL is not set", http.StatusInternalServerError)
		return
	}

	// Call the external service
	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, "Failed to call service: " + err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}

	// Write status code
	w.WriteHeader(resp.StatusCode)

	// Copy response body
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("Error copying response body: %v", err)
	}
}

// getEnv reads an environment variable or returns a fallback value.
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

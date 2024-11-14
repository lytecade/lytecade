package main

import (
    "fmt"
    "net/http"
)

// Handler function to serve HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "LYTECADE")
}

func main() {
    // Set up route and handler
    http.HandleFunc("/", handler)

    // Start the web server
    fmt.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}


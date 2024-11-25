package main

import (
	"fmt"
	"net/http"
	"github.com/lytecade/lytecade/internal/handlers"
)


func main() {
	http.HandleFunc("/", handlers.RunIndexHandler)
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


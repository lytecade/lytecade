package handlers

import (
	"fmt"
	"net/http"
)

func RunIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LYTECADE-HANDLER")
}

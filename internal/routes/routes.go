package routes

import (
	"fmt"
	"net/http"
	"github.com/lytecade/lytecade/internal/pages"
)

func RoutesInit() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", pages.HandleHomePage)
	http.HandleFunc("/about", pages.HandleAboutPage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}


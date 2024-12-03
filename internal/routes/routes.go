package routes

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"github.com/lytecade/lytecade/internal/pages"
	"path/filepath"
)

var (
    currentSites []string
)

func RouteInit() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))
	http.HandleFunc("/", pages.HandleHomePage)
	http.HandleFunc("/about", pages.HandleAboutPage)
}

func RouteSites() {
	err := filepath.Walk("./web/sites", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		fmt.Println(path)
		currentSite := strings.TrimPrefix(path, "./web/sites")
		if currentSite == "" {
			return nil
		}
        if !contains(currentSites, currentSite) {
            currentSites = append(currentSites, currentSite)
        }
		http.HandleFunc(fmt.Sprintf("/%s", currentSite), func(w http.ResponseWriter, r *http.Request) {
			pages.HandleGamePage(w, r, currentSite)
		})
		http.Handle("/" + currentSite + "/assets/", pages.HandleGameFolder(currentSite, path, "assets"))
		http.Handle("/" + currentSite + "/src/", pages.HandleGameFolder(currentSite, path, "src"))

		return nil
	})
	if err != nil {
		fmt.Println("Error registering game server:", err)
	}
}

func RouteListen() {
	fmt.Println("Listening on Port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func GetRouteSites() []string {
    return append([]string(nil), currentSites...)
}

func contains(setSites []string, siteValue string) bool {
    for _, currentChar := range setSites {
        if currentChar == siteValue {
            return true
        }
    }
    return false
}




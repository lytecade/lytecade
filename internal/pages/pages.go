package pages

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title          string
	Description    string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplFiles := []string{
		"web/templates/layout.tmpl", 
		tmpl,          
	}
	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		http.Error(w, "Error parsing templates: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func HandleHomePage(w http.ResponseWriter, r *http.Request) {
	data := PageData {
		Title:          "Home - LyteCade",
		Description:    "This is the homepage for LyteCade.",
	}
	renderTemplate(w, "web/templates/index.tmpl", data)
}

func HandleAboutPage(w http.ResponseWriter, r *http.Request) {
	data := PageData {
		Title:          "About - LyteCade",
		Description:    "This is the about page for LyteCade.",
	}
	renderTemplate(w, "web/templates/about.tmpl", data)
}

func HandleGamePage(w http.ResponseWriter, r *http.Request, gameName string) {
	sitePath := filepath.Join("./web/sites", gameName, "index.html")
	fmt.Println("Files for Game:", sitePath)
}

func HandleGameFolder(siteName, originalPath, resourceFolder string) http.Handler {
	filePrefix := ("/" + siteName + "/" + resourceFolder)
	fileServer := http.StripPrefix(filePrefix, http.FileServer(http.Dir(filepath.Join(originalPath, resourceFolder))))
	return fileServer
}

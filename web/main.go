package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title        string
	Message      template.HTML
	ErrorCode    int
	ErrorMessage string
}

func renderTemplate(w http.ResponseWriter, tmplFile string, data PageData) {
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{
		Title:   "Home",
		Message: template.HTML("<b>Bienvenido</b> a la página de inicio"),
	}

	tmplFile := filepath.Join("web/templates/", "index.html")
	renderTemplate(w, tmplFile, pageData)

}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{
		Title:        "Error 404",
		ErrorCode:    404,
		ErrorMessage: "¡Página no encontrada!",
	}

	tmplFile := filepath.Join("web/templates/", "error.html")
	renderTemplate(w, tmplFile, pageData)

}

func main() {
	fs := http.FileServer(http.Dir(filepath.Join("web", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/error", errorHandler)

	fmt.Println("El servidor está escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

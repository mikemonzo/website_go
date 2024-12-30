package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikemonzo/website_go/internal/models"
)

func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pageData := models.PageData{
		Title:   "Home",
		Author:  "Mike Monzo",
		Welcome: "Bienvenido a la página de inicio",
	}

	page := r.URL.Path[1:]

	if page != "" {
		page = "index.html"
	}

	tmplFile := "web/templates/" + page

	if _, err := os.Stat(tmplFile); err != nil {
		tmplFile = "web/templates/error.html"

		pageData.ErrorCode = http.StatusNotFound
		pageData.ErrorMessage = "¡Página no encontrada!"
	}

	renderTemplate(w, tmplFile, pageData)

}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	pageData := models.PageData{
		Title:        "¡Página no encontrada!",
		ErrorCode:    http.StatusInternalServerError,
		ErrorMessage: "Error interno del servidor",
	}

	tmplFile := filepath.Join("web/templates/", "error.html")
	renderTemplate(w, tmplFile, pageData)

}

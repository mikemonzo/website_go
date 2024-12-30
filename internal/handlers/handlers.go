package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikemonzo/website_go/internal/models"
)

func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
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

	if page == "" {
		page = "index.html"
	}

	headContent, err := os.ReadFile("web/templates/header.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	headTemplate, err := template.New("header").Parse(string(headContent))
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var headBuffer bytes.Buffer
	err = headTemplate.Execute(&headBuffer, pageData)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	pageData.HeaderContent = template.HTML(headBuffer.String())

	navbarContent, err := os.ReadFile("web/templates/navbar.html")
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	navbarTemplate, err := template.New("navbar").Parse(string(navbarContent))
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var navbarBuffer bytes.Buffer
	err = navbarTemplate.Execute(&navbarBuffer, pageData)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	pageData.NavbarContent = template.HTML(navbarBuffer.String())

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

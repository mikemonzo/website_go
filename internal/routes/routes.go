package routes

import (
	"net/http"

	"github.com/mikemonzo/website_go/internal/handlers"
)

func LoadRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

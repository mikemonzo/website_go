package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mikemonzo/website_go/internal/config"
	"github.com/mikemonzo/website_go/internal/routes"
)

func main() {

	cfg := config.LoadConfig()

	routes.LoadRoutes()

	addr := cfg.HTTPHOST + ":" + cfg.HTTPPort

	server := &http.Server{Addr: addr, Handler: nil}

	fmt.Println("Servidor web escuchando en http://" + addr)
	log.Fatal(server.ListenAndServe())
}

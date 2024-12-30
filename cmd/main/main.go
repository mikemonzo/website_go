package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mikemonzo/website_go/internal/config"
	"github.com/mikemonzo/website_go/internal/routes"
)

func main() {

	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error al abrir el archivo de log")
		os.Exit(1)
	}

	defer logFile.Close()

	log.SetOutput(logFile)

	cfg := config.LoadConfig()

	routes.LoadRoutes()

	addr := cfg.HTTPHost + ":" + cfg.HTTPPort

	server := &http.Server{Addr: addr, Handler: nil}

	fmt.Println("Servidor web escuchando en http://" + addr)
	log.Fatal(server.ListenAndServe())
}

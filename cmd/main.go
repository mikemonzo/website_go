package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/about", abouthandler)
	http.HandleFunc("/contact", contacthandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la página de inicio!")
}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Esta es la página 'Acerca de nosotros'")
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Póngase en contacto con nosotros en mikemonzo.com")
}

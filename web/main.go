package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/detail/", productDetailHandler)

	fmt.Println("El servidor está escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Bienvenido a la página de inicio!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aquí se muestran todos los productos disponibles")
}

func productDetailHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Path[len("/products/detail/"):]
	fmt.Fprintf(w, "Detalles del producto con ID: %s", productID)
}

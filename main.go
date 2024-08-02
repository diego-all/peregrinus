package main

import (
	"log"
	"net/http"
)

func main() {
	// Directorio de archivos estáticos
	directorio := "./sitio"

	// static files
	fileServer := http.FileServer(http.Dir(directorio))

	// Registrar el manejador de archivos en el servidor HTTP
	http.Handle("/sitio/", http.StripPrefix("/sitio", fileServer))

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor de archivos escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

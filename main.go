package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Report struct {
	Command string `json:"command"`
	Output  string `json:"output"`
}

func reportHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var report Report

	if err := json.Unmarshal(body, &report); err != nil {
		http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Recibido reporte: %v\n", report)

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Reporte recibido correctamente"))

}

func main() {
	// Directorio de archivos estáticos
	directorio := "./sitio"

	// static files
	fileServer := http.FileServer(http.Dir(directorio))

	// Registrar el manejador de archivos en el servidor HTTP
	http.Handle("/sitio/", http.StripPrefix("/sitio", fileServer))
	http.HandleFunc("/report", reportHandler)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor de archivos escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

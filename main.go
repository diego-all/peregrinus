package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	// Directorio de archivos estáticos
	directorio := "./public"

	// static files
	fileServer := http.FileServer(http.Dir(directorio))

	// Registrar el manejador de archivos en el servidor HTTP
	http.Handle("/public/", http.StripPrefix("/public", fileServer))

	// Endpoint para ejecutar comandos
	http.HandleFunc("/execute", executeCommand)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor de archivos escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func executeCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "No se pudo leer el cuerpo de la solicitud", http.StatusInternalServerError)
		return
	}

	cmdStr := string(body)
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al ejecutar el comando: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

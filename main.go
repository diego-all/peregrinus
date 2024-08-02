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

	// Endpoint para ejecutar comandos en el contexto del sistema operativo
	http.HandleFunc("/execute-system", executeCommandSystem)

	// Endpoint para ejecutar comandos en el contexto del servidor de archivos
	http.HandleFunc("/execute-files", executeCommandFiles)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor de archivos escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Endpoint para ejecutar comandos en el contexto del sistema operativo
func executeCommandSystem(w http.ResponseWriter, r *http.Request) {
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

// Endpoint para ejecutar comandos en el contexto del servidor de archivos
func executeCommandFiles(w http.ResponseWriter, r *http.Request) {
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
	cmd := exec.Command("bash", "-c", fmt.Sprintf("cd public && %s", cmdStr))
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al ejecutar el comando: %s", err), http.StatusInternalServerError)
		return
	}

	w.Write(output)
}

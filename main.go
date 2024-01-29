package main

import (
	"fmt"
	"net/http"
	v1 "newsletter-app/pkg/api/v1"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
)

func main() {
	// Conectar a MongoDB
	err := mongodb.Connect("mongodb+srv://gustavohdzmz:COERlJXgVI3XSp6M@newsletter.9soh00l.mongodb.net/?retryWrites=true&w=majority")
	if err != nil {
		fmt.Println("Error al conectar a MongoDB:", err)
		return
	} else {
		fmt.Println("Conexión a la base de datos exitosa")
	}
	defer mongodb.Disconnect()

	// Configurar el enrutador con el servicio de suscriptores y el sender de correo electrónico
	router := v1.SetupRouter()

	// Iniciar el servidor
	port := 8080
	address := fmt.Sprintf(":%d", port)

	fmt.Printf("Servidor escuchando en http://localhost%s\n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		fmt.Println(err)
	}
}

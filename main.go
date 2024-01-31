package main

import (
	"fmt"
	"net/http"
	v1 "newsletter-app/pkg/api/v1"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"os"

	"github.com/gorilla/handlers"

	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables de entorno desde el archivo .env
	env := godotenv.Load()
	if env != nil {
		fmt.Printf("Error loading .env file: %v\n", env)
		return
	}
	// Conectar a MongoDB
	mongoUrl := os.Getenv("mongoUrl")
	fmt.Println("MongoDB:", mongoUrl)
	err := mongodb.Connect(mongoUrl)
	if err != nil {
		fmt.Println("Error al conectar a MongoDB:", err)
		return
	} else {
		fmt.Println("Conexión a la base de datos exitosa")
	}
	defer mongodb.Disconnect()

	// Configurar el enrutador con el servicio de suscriptores y el sender de correo electrónico
	router := v1.SetupRouter()

	// Habilitar CORS permitiendo cualquier origen
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	router.Use(handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders))

	// Agregar manejador OPTIONS al enrutador
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Iniciar el servidor
	port := 8080
	address := fmt.Sprintf(":%d", port)

	fmt.Printf("Servidor escuchando en http://localhost%s/docs/index.html\n", address)
	err = http.ListenAndServe(address, router)
	if err != nil {
		fmt.Println(err)
	}
}

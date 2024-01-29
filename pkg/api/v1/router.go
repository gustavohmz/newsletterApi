package v1

import (
	"newsletter-app/pkg/service"

	"newsletter-app/pkg/infrastructure/adapters/email"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Newsletter API
// @version 1.0
// @description API para gestionar suscriptores y boletines del boletín.
// @termsOfService http://swagger.io/terms/
// @contact name@newsletter.com
// @license MIT
// @host localhost:8080
// @BasePath /api/v1
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Crear una instancia del servicio de suscriptores
	subscriberService := service.NewSubscriberService()
	newsletterService := service.NewNewsletterService()

	// Crear una instancia del sender de correo electrónico
	emailSender := email.NewBrevoEmailSender()

	// Configuración de rutas
	r.HandleFunc("/api/v1/subscribe/{email}", SubscribeHandler(subscriberService)).Methods("POST")
	r.HandleFunc("/api/v1/unsubscribe/{email}", UnsubscribeHandler(subscriberService)).Methods("DELETE")
	r.HandleFunc("/api/v1/subscribers/{email}", GetSubscriberHandler(subscriberService)).Methods("GET")

	// Agregar ruta para enviar boletín
	r.HandleFunc("/api/v1/newsletters/send/{newsletterID}", SendNewsletterHandler(subscriberService, newsletterService, emailSender)).Methods("POST")
	r.HandleFunc("/api/v1/newsletters", CreateNewsletterHandler(newsletterService)).Methods("POST")

	// Ruta para Swagger
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	return r
}

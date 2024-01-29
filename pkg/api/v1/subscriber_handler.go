// pkg/api/v1/subscriber_handler.go

package v1

import (
	"encoding/json"
	"net/http"
	_ "newsletter-app/docs"
	"newsletter-app/pkg/service"
	"regexp"

	"github.com/gorilla/mux"
)

// ErrorResponse estructura para respuestas de error
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Summary Subscribe to the newsletter
// @Description Allows a user to subscribe to the newsletter
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email path string true "Email address to subscribe"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/subscribe/{email} [post]
func SubscribeHandler(subscriberService *service.SubscriberService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		if email == "" || !isValidEmail(email) {
			respondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		// Verificar si el correo ya está suscrito
		existingSubscriber, err := subscriberService.GetSubscriberByEmail(email)
		if err == nil && existingSubscriber != nil {
			respondWithError(w, http.StatusConflict, "User is already subscribed")
			return
		}

		// Llamar al servicio para realizar la suscripción
		err = subscriberService.Subscribe(email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to subscribe user")
			return
		}

		// Obtener información detallada del suscriptor
		subscriber, err := subscriberService.GetSubscriberByEmail(email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to get subscriber details")
			return
		}

		// Responder con éxito y la información detallada del suscriptor
		respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":     "OK",
			"message":    "subscription successful",
			"subscriber": subscriber,
		})
	}
}

// @Summary Unsubscribe from the newsletter
// @Description Allows a user to unsubscribe from the newsletter
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email path string true "Email address to unsubscribe"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/unsubscribe/{email} [delete]
func UnsubscribeHandler(subscriberService *service.SubscriberService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		if email == "" || !isValidEmail(email) {
			respondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		// Llamar al servicio para cancelar la suscripción
		err := subscriberService.Unsubscribe(email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to unsubscribe user")
			return
		}

		// Responder con éxito
		respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "User unsubscribed successfully",
		})
	}
}

// @Summary Get subscriber by email
// @Description Get details of a subscriber by email address
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email path string true "Email address to get details"
// @Success 200 {object} domain.Subscriber
// @Failure 404 {object} ErrorResponse "Subscriber not found"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/subscribers/{email} [get]
func GetSubscriberHandler(subscriberService *service.SubscriberService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		if email == "" || !isValidEmail(email) {
			respondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		// Llamar al servicio para obtener el suscriptor por email
		subscriber, err := subscriberService.GetSubscriberByEmail(email)
		if err != nil {
			// Verificar si el suscriptor no fue encontrado
			if err.Error() == "mongo: no documents in result" {
				respondWithError(w, http.StatusNotFound, "Subscriber not found")
				return
			}

			respondWithError(w, http.StatusInternalServerError, "Failed to get subscriber")
			return
		}

		// Llamar al servicio para eliminar el suscriptor
		err = subscriberService.Unsubscribe(email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to unsubscribe user")
			return
		}

		// Responder con el objeto Subscriber en formato JSON
		respondWithJSON(w, http.StatusOK, subscriber)
	}
}

// Función de utilidad para responder con un error en formato JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON envía una respuesta JSON al cliente.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Función para validar el formato del correo electrónico
func isValidEmail(email string) bool {
	// Utiliza una expresión regular simple para validar el formato del correo
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

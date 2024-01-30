// pkg/api/v1/subscriber_handler.go

package v1

import (
	"net/http"
	_ "newsletter-app/docs"
	"newsletter-app/pkg/service"

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
// @Param category path string true "Category to subscribe to"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/subscribe/{email}/{category} [post]
func SubscribeHandler(subscriberService *service.SubscriberService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		category := mux.Vars(r)["category"]
		if email == "" || !service.IsValidEmail(email) {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		// Verificar si el correo ya está suscrito
		existingSubscriber, err := subscriberService.GetSubscriberByEmail(email)
		if err == nil && existingSubscriber != nil {
			service.RespondWithError(w, http.StatusConflict, "User is already subscribed")
			return
		}

		// Llamar al servicio para realizar la suscripción
		err = subscriberService.Subscribe(email, category)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to subscribe user")
			return
		}

		// Obtener información detallada del suscriptor
		subscriber, err := subscriberService.GetSubscriberByEmail(email)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to get subscriber details")
			return
		}

		// Responder con éxito y la información detallada del suscriptor
		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
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
		if email == "" || !service.IsValidEmail(email) {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		// Llamar al servicio para cancelar la suscripción
		err := subscriberService.Unsubscribe(email)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to unsubscribe user")
			return
		}

		// Responder con éxito
		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
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
		if email == "" || !service.IsValidEmail(email) {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		// Llamar al servicio para obtener el suscriptor por email
		subscriber, err := subscriberService.GetSubscriberByEmail(email)
		if err != nil {
			// Verificar si el suscriptor no fue encontrado
			if err.Error() == "mongo: no documents in result" {
				service.RespondWithError(w, http.StatusNotFound, "Subscriber not found")
				return
			}

			service.RespondWithError(w, http.StatusInternalServerError, "Failed to get subscriber")
			return
		}

		// Llamar al servicio para eliminar el suscriptor
		err = subscriberService.Unsubscribe(email)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to unsubscribe user")
			return
		}

		// Responder con el objeto Subscriber en formato JSON
		service.RespondWithJSON(w, http.StatusOK, subscriber)
	}
}

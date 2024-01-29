// pkg/api/v1/newsletter_handler.go

package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service"

	"bytes"

	"github.com/gorilla/mux"
)

// @Summary Send newsletter to subscribers
// @Description Allows an admin user to send a newsletter to a list of subscribers
// @Tags newsletters
// @Accept json
// @Produce json
// @Param newsletterID path string true "ID of the newsletter to be sent"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/newsletters/send/{newsletterID} [post]
func SendNewsletterHandler(subscriberService *service.SubscriberService, newsletterService *service.NewsletterService, emailSender email.Sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del boletín desde los parámetros de la ruta
		newsletterID := mux.Vars(r)["newsletterID"]
		if newsletterID == "" {
			respondWithError(w, http.StatusBadRequest, "Invalid newsletter ID")
			return
		}

		// Obtener el boletín por ID
		newsletter, err := newsletterService.GetNewsletterByID(newsletterID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve newsletter")
			return
		}

		// Verificar si el boletín ha sido enviado previamente
		if newsletter.Sent != nil && *newsletter.Sent {
			respondWithError(w, http.StatusBadRequest, "Newsletter has already been sent")
			return
		}

		// Obtener la lista de destinatarios del boletín
		recipients := newsletter.SentRecipients

		// Verificar si hay destinatarios para enviar el boletín
		if len(recipients) == 0 {
			respondWithError(w, http.StatusBadRequest, "No recipients to send the newsletter to")
			return
		}

		// Obtener el contenido del boletín del cuerpo de la solicitud
		var requestBody struct {
			Content string `json:"content"`
		}

		// Asignar un contenido de prueba (puedes personalizarlo según tus necesidades)
		requestBody.Content = "Contenido de prueba del boletín"

		// Convertir el contenido a bytes para simular la decodificación desde el cuerpo JSON
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to marshal test content")
			return
		}

		// Decodificar el cuerpo JSON
		// err = json.NewDecoder(r.Body).Decode(&requestBody)
		// if err != nil {
		// 	respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		// 	return
		// }

		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&requestBody)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Iterar sobre los destinatarios y enviar el boletín
		for _, recipient := range recipients {
			// Llamar al servicio para obtener el suscriptor por email
			subscriber, err := subscriberService.GetSubscriberByEmail(recipient.Email)
			if err != nil {
				// Verificar si hubo un error diferente al no encontrar documentos
				if err.Error() != "mongo: no documents in result" {
					respondWithError(w, http.StatusInternalServerError, "Failed to get subscriber")
					return
				}
				// Si no se encuentra el suscriptor, continuar con el siguiente destinatario
				fmt.Printf("Subscriber not found for email %s\n", recipient.Email)
				continue
			}

			// Imprimir información del suscriptor (para depuración)
			fmt.Printf("Subscriber: %+v\n", subscriber)

			// Utilizar el contenido decodificado en cada iteración
			newsletterContent := requestBody.Content

			// Enviar boletín al destinatario
			err = emailSender.Send("Asunto del Boletín", newsletterContent, []string{recipient.Email})
			if err != nil {
				// Manejar el error (puedes logearlo, enviar una respuesta específica, etc.)
				fmt.Printf("Error sending newsletter to %s: %s\n", recipient.Email, err.Error())
				continue
			}

			// Registrar el destinatario al que se le envió el boletín
			fmt.Printf("Newsletter sent to %s\n", recipient.Email)
		}

		// Actualizar el boletín como enviado
		err = newsletterService.UpdateNewsletter(newsletterID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to update newsletter")
			return
		}
		// Responder con éxito
		respondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter sent successfully",
		})
	}
}

// @Summary Create a new newsletter
// @Description Allows an admin user to create a new newsletter
// @Tags newsletters
// @Accept json
// @Produce json
// @Param newsletter body domain.Newsletter true "Newsletter details"
// @Success 201 {string} string "Created"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/newsletters [post]
func CreateNewsletterHandler(newsletterService *service.NewsletterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud en una estructura de Newsletter
		var newNewsletter domain.Newsletter
		err := json.NewDecoder(r.Body).Decode(&newNewsletter)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Validar que se proporcionen categorías
		if len(newNewsletter.Categories) == 0 {
			respondWithError(w, http.StatusBadRequest, "Categories are required")
			return
		}

		// Validar que las categorías proporcionadas sean válidas
		for _, cat := range newNewsletter.Categories {
			if !isValidCategory(cat) {
				respondWithError(w, http.StatusBadRequest, "Invalid category specified")
				return
			}
		}

		// Lógica para crear el nuevo boletín
		err = newsletterService.SaveNewsletter(newNewsletter)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Failed to create newsletter")
			return
		}

		// Responder con éxito
		respondWithJSON(w, http.StatusCreated, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter created successfully",
			"sent":    newNewsletter.Sent,
		})
	}
}

// isValidCategory verifica si la categoría proporcionada es válida.
func isValidCategory(category domain.Category) bool {
	switch category {
	case domain.SpecialOffers, domain.Memberships, domain.MonthlyPromotions, domain.NewProducts:
		return true
	default:
		return false
	}
}

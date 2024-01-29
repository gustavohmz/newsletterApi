// pkg/api/v1/newsletter_handler.go

package v1

import (
	"encoding/json"
	"net/http"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service"

	"github.com/gorilla/mux"
)

// @Summary Send newsletter to subscribers
// @Description Allows an admin user to send a newsletter to a list of subscribers
// @Tags newsletters
// @Accept json
// @Produce json
// @Param newsletterID path string true "ID of the newsletter to be sent"
// @Param requestBody body SendNewsletterRequest true "Request body containing recipients"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/newsletters/send/{newsletterID} [post]
func SendNewsletterHandler(subscriberService *service.SubscriberService, newsletterService *service.NewsletterService, emailSender email.Sender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del boletín desde los parámetros de la ruta
		newsletterID := mux.Vars(r)["newsletterID"]
		if newsletterID == "" {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid newsletter ID")
			return
		}

	}
}

// Estructura para la solicitud de envío de boletín
type SendNewsletterRequest struct {
	Recipients []struct {
		Email string `json:"email"`
	} `json:"recipients"`
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
			service.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Validar que se proporcionen categorías
		if len(newNewsletter.Categories) == 0 {
			service.RespondWithError(w, http.StatusBadRequest, "Categories are required")
			return
		}

		// Validar que las categorías proporcionadas sean válidas
		for _, cat := range newNewsletter.Categories {
			if !isValidCategory(cat) {
				service.RespondWithError(w, http.StatusBadRequest, "Invalid category specified")
				return
			}
		}

		// Lógica para crear el nuevo boletín
		err = newsletterService.SaveNewsletter(newNewsletter)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to create newsletter")
			return
		}

		// Responder con éxito
		service.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter created successfully",
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

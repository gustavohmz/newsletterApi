// pkg/api/v1/newsletter_handler.go

package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service"
	"strconv"

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
			service.RespondWithError(w, http.StatusBadRequest, "Invalid newsletter ID")
			return
		}

		// Llamar a la función SendNewsletter en el servicio
		err := newsletterService.SendNewsletter(w, r, newsletterID, emailSender)
		if err != nil {
			// Manejar el error (puedes logearlo, enviar una respuesta específica, etc.)
			fmt.Printf("Error sending newsletter: %s\n", err.Error())
			return
		}

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

// @Summary Get a list of newsletters
// @Description Retrieves a list of newsletters with optional search and pagination parameters
// @Tags newsletters
// @Accept json
// @Produce json
// @Param name query string false "Name of the newsletter to search for"
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page for pagination"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/newsletters [get]
func GetNewslettersHandler(newsletterService *service.NewsletterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener parámetros de consulta
		name := r.URL.Query().Get("name")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

		// Llamar a la función en el servicio para obtener la lista de boletines
		newsletters, err := newsletterService.GetNewsletters(name, page, pageSize)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve newsletters")
			return
		}

		// Responder con la lista de boletines
		service.RespondWithJSON(w, http.StatusOK, newsletters)
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

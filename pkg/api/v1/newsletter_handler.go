package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service"
	"newsletter-app/pkg/service/Dtos/request"
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

		// Validar que se proporcione la categoría
		if newNewsletter.Category == "" {
			service.RespondWithError(w, http.StatusBadRequest, "Category is required")
			return
		}

		existingNewsletter, err := newsletterService.GetNewsletterByCategory(newNewsletter.Category)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to check for existing newsletters")
			return
		}

		// Si ya existe un boletín con la misma categoría, retornar un error de conflicto
		if existingNewsletter != nil {
			service.RespondWithError(w, http.StatusConflict, "Newsletter with the specified category already exists")
			return
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

// @Summary Update an existing newsletter
// @Description Allows an admin user to update an existing newsletter
// @Tags newsletters
// @Accept json
// @Produce json
// @Param updateRequest body request.UpdateNewsletterRequest true "Update newsletter details"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/newsletters [put]
func UpdateNewsletterHandler(newsletterService *service.NewsletterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decodificar el cuerpo de la solicitud en una estructura de UpdateNewsletterRequest
		var updateRequest request.UpdateNewsletterRequest
		err := json.NewDecoder(r.Body).Decode(&updateRequest)
		if err != nil {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		// Lógica para actualizar el boletín
		err = newsletterService.UpdateNewsletter(updateRequest)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to update newsletter")
			return
		}

		// Responder con éxito
		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter updated successfully",
		})
	}
}

// @Summary Delete a newsletter
// @Description Allows an admin user to delete a newsletter
// @Tags newsletters
// @Accept json
// @Produce json
// @Param id path string true "ID of the newsletter to delete"
// @Success 200 {string} string "OK"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /api/v1/newsletters/{id} [delete]
func DeleteNewsletterHandler(newsletterService *service.NewsletterService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del boletín de los parámetros de la ruta
		id := mux.Vars(r)["id"]
		if id == "" {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing newsletter ID")
			return
		}

		// Llamar al servicio para eliminar el boletín
		err := newsletterService.DeleteNewsletter(id)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to delete newsletter")
			return
		}

		// Responder con éxito
		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter deleted successfully",
		})
	}
}

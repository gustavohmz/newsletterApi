// pkg/service/newsletter_service.go

package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"time"
)

// NewsletterService es una estructura que maneja la lógica de negocio relacionada con los boletines.
type NewsletterService struct {
	newsletterRepository *mongodb.NewsletterRepository
	subscriberRepository *mongodb.SubscriberRepository
}

// NewNewsletterService crea una nueva instancia de NewsletterService.
func NewNewsletterService() *NewsletterService {
	return &NewsletterService{
		newsletterRepository: mongodb.NewNewsletterRepository(),
		subscriberRepository: mongodb.NewSubscriberRepository(),
	}
}

// Función principal para crear boletines
func (s *NewsletterService) SaveNewsletter(newsletter domain.Newsletter) error {
	var decodedAttachments []domain.Attachment

	// Iterar sobre los archivos adjuntos en base64 y procesarlos
	for _, base64Attachment := range newsletter.Attachments {
		attachment := domain.Attachment{
			Name: base64Attachment.Name,
			Data: base64Attachment.Data,
		}
		decodedAttachments = append(decodedAttachments, attachment)
	}

	// Reemplazar el slice original con los adjuntos decodificados
	newsletter.Attachments = decodedAttachments

	// Lógica para guardar el boletín
	return s.newsletterRepository.SaveNewsletter(newsletter)
}

// UpdateNewsletter actualiza el campo "sent" del boletín.
func (s *NewsletterService) UpdateNewsletter(newsletterID string) error {
	return s.newsletterRepository.UpdateNewsletter(newsletterID)
}

// GetNewsletterByID obtiene un boletín por ID.
func (s *NewsletterService) GetNewsletterByID(newsletterID string) (*domain.Newsletter, error) {
	return s.newsletterRepository.GetNewsletterByID(newsletterID)
}

// GetNewsletters obtiene una lista de boletines con opciones de búsqueda y paginación.
func (s *NewsletterService) GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error) {
	newsletters, err := s.newsletterRepository.GetNewsletters(searchName, page, pageSize)
	if err != nil {
		return nil, err
	}

	return newsletters, nil
}

// EnvíarNewsletter envía un boletín a una lista de suscriptores.
func (s *NewsletterService) SendNewsletter(w http.ResponseWriter, r *http.Request, newsletterID string, emailSender email.Sender) error {
	// Obtener el boletín por ID
	newsletter, err := s.GetNewsletterByID(newsletterID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve newsletter")
		return err
	}

	// Obtener la lista de destinatarios del cuerpo de la solicitud
	var requestBody SendNewsletterRequest
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return err
	}

	// Verificar si hay destinatarios para enviar el boletín
	if len(requestBody.Recipients) == 0 {
		RespondWithError(w, http.StatusBadRequest, "No recipients to send the newsletter to")
		return nil
	}

	// Iterar sobre los destinatarios y enviar el boletín
	for _, recipient := range requestBody.Recipients {
		// Llamar al servicio para obtener el suscriptor por email
		subscriber, err := s.subscriberRepository.GetSubscriberByEmail(recipient.Email)
		if err != nil {
			// Verificar si hubo un error diferente al no encontrar documentos
			if err.Error() != "mongo: no documents in result" {
				RespondWithError(w, http.StatusInternalServerError, "Failed to get subscriber")
				return err
			}
			// Si no se encuentra el suscriptor, continuar con el siguiente destinatario
			fmt.Printf("Subscriber not found for email %s\n", recipient.Email)
			continue
		}

		// Imprimir información del suscriptor (para depuración)
		fmt.Printf("Subscriber: %+v\n", subscriber)

		// Verificar si hay contenido en el boletín
		if newsletter.Content == "" {
			RespondWithError(w, http.StatusBadRequest, "Newsletter content is empty")
			return nil
		}

		// Enviar boletín al destinatario
		err = emailSender.Send("Asunto del Boletín", newsletter.Content, []string{recipient.Email})
		if err != nil {
			// Manejar el error (puedes logearlo, enviar una respuesta específica, etc.)
			fmt.Printf("Error sending newsletter to %s: %s\n", recipient.Email, err.Error())
			continue
		}

		// Registrar el destinatario al que se le envió el boletín
		fmt.Printf("Newsletter sent to %s\n", recipient.Email)
	}

	// Responder con éxito
	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "Newsletter sent successfully",
	})
	return nil
}

// Estructura para la solicitud de envío de boletín
type SendNewsletterRequest struct {
	Recipients []struct {
		Email string `json:"email"`
	} `json:"recipients"`
}

// ProgramarEnvío programa el envío de un boletín en una fecha y hora específicas.
func (s *NewsletterService) ProgramarEnvío(newsletterID string, scheduleTime time.Time) error {
	// Lógica para programar el envío del boletín
	return nil
}

func decodeBase64Attachment(name, data string) (domain.Attachment, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return domain.Attachment{}, err
	}

	return domain.Attachment{
		Name: name,
		Data: string(decodedData),
	}, nil
}

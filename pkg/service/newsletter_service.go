// pkg/service/newsletter_service.go

package service

import (
	"encoding/base64"
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
			Type: base64Attachment.Type,
		}
		decodedAttachments = append(decodedAttachments, attachment)
	}

	// Reemplazar el slice original con los adjuntos decodificados
	newsletter.Attachments = decodedAttachments
	return s.newsletterRepository.SaveNewsletter(newsletter)
}

func (s *NewsletterService) GetNewsletterByCategory(category string) (*domain.Newsletter, error) {
	return s.newsletterRepository.GetNewsletterByCategory(category)
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

	// Obtener la lista de suscriptores con la misma categoría del boletín
	subscribers, err := s.subscriberRepository.GetSubscribersByCategory(newsletter.Category)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve subscribers")
		return err
	}

	// Verificar si hay suscriptores para enviar el boletín
	if len(subscribers) == 0 {
		RespondWithError(w, http.StatusBadRequest, "No subscribers to send the newsletter to")
		return nil
	}

	// Iterar sobre los suscriptores y enviar el boletín
	for _, subscriber := range subscribers {
		// Imprimir información del suscriptor (para depuración)
		fmt.Printf("Subscriber: %+v\n", subscriber)

		// Verificar si hay contenido en el boletín
		if newsletter.Content == "" {
			RespondWithError(w, http.StatusBadRequest, "Newsletter content is empty")
			return nil
		}

		// Enviar boletín al suscriptor
		err = emailSender.Send(newsletter.Subject, newsletter.Content, []string{subscriber.Email})
		if err != nil {
			// Manejar el error (puedes logearlo, enviar una respuesta específica, etc.)
			fmt.Printf("Error sending newsletter to %s: %s\n", subscriber.Email, err.Error())
			continue
		}

		// Registrar el suscriptor al que se le envió el boletín
		fmt.Printf("Newsletter sent to %s\n", subscriber.Email)
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

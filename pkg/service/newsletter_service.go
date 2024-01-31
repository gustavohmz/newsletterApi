package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"newsletter-app/pkg/service/Dtos/request"
	"strings"
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

	// Decodificar los archivos adjuntos del boletín
	decodedAttachments, err := DecodeAttachments(newsletter.Attachments)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to decode attachments")
		return err
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

		// Reemplazar variables en el contenido del boletín
		emailCategoryConcatenation := fmt.Sprintf("%s|%s", subscriber.Email, subscriber.Category)
		newsletterContent := strings.ReplaceAll(newsletter.Content, "{email}", emailCategoryConcatenation)

		// Enviar boletín al suscriptor con archivos adjuntos
		err = emailSender.Send(newsletter.Subject, newsletterContent, []string{subscriber.Email}, decodedAttachments)
		if err != nil {
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

// DecodeAttachments decodifica un conjunto de objetos Attachment y retorna los datos de los archivos adjuntos.
func DecodeAttachments(attachments []domain.Attachment) ([]*domain.Attachment, error) {
	var decodedAttachments []*domain.Attachment

	for _, attachment := range attachments {
		_, err := base64.StdEncoding.DecodeString(attachment.Data)
		if err != nil {
			return nil, errors.New("failed to decode attachment data")
		}

		decodedAttachment := &domain.Attachment{
			Name: attachment.Name,
			Data: attachment.Data,
			Type: attachment.Type,
		}

		decodedAttachments = append(decodedAttachments, decodedAttachment)
	}

	return decodedAttachments, nil
}

// UpdateNewsletter actualiza un boletín existente.
func (s *NewsletterService) UpdateNewsletter(updateRequest request.UpdateNewsletterRequest) error {
	// Validar que se proporcione el ID
	if updateRequest.ID.IsZero() {
		return errors.New("ID is required for update")
	}

	// Obtener el boletín existente por ID
	existingNewsletter, err := s.GetNewsletterByID(updateRequest.ID.Hex())
	if err != nil {
		return err
	}

	// Actualizar los campos necesarios
	existingNewsletter.Name = updateRequest.Name
	existingNewsletter.Category = updateRequest.Category
	existingNewsletter.Subject = updateRequest.Subject
	existingNewsletter.Content = updateRequest.Content

	// Actualizar los archivos adjuntos si se proporcionan en la solicitud de actualización
	if len(updateRequest.Attachments) > 0 {
		existingNewsletter.Attachments = make([]domain.Attachment, len(updateRequest.Attachments))
		for i, attachment := range updateRequest.Attachments {
			existingNewsletter.Attachments[i] = domain.Attachment{
				Name: attachment.Name,
				Data: attachment.Data,
				Type: attachment.Type,
			}
		}
	} else {
		// No se proporcionan nuevos archivos adjuntos, eliminar los existentes
		existingNewsletter.Attachments = nil
	}

	// Llamar a la función de actualización en el repositorio
	return s.newsletterRepository.UpdateNewsletter(*existingNewsletter)
}

// DeleteNewsletter elimina un boletín por su ID.
func (s *NewsletterService) DeleteNewsletter(id string) error {
	return s.newsletterRepository.DeleteNewsletterByID(id)
}

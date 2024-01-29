// pkg/service/newsletter_service.go

package service

import (
	"encoding/base64"
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"time"
)

// NewsletterService es una estructura que maneja la lógica de negocio relacionada con los boletines.
type NewsletterService struct {
	newsletterRepository *mongodb.NewsletterRepository
}

// NewNewsletterService crea una nueva instancia de NewsletterService.
func NewNewsletterService() *NewsletterService {
	return &NewsletterService{
		newsletterRepository: mongodb.NewNewsletterRepository(),
	}
}

// Función principal para crear boletines
func (s *NewsletterService) SaveNewsletter(newsletter domain.Newsletter) error {
	// Crear un nuevo slice para almacenar los adjuntos decodificados
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

// EnvíarNewsletter envía un boletín a una lista de suscriptores.
func (s *NewsletterService) SendNewsletter(newsletterID string, subscribers []domain.Subscriber) error {
	// Lógica para enviar el boletín a la lista de suscriptores
	// Puedes utilizar un servicio externo para enviar correos electrónicos, como SendGrid o SMTP
	return nil
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

// isValidFileType verifica si el tipo de archivo es válido.
func isValidFileType(data []byte) bool {
	// Implementa la lógica para validar el tipo de archivo
	// Puedes usar bibliotecas o realizar tus propias verificaciones
	// En este ejemplo, simplemente se asume que es un PDF, pero debes mejorar esta lógica.
	// Por ejemplo, podrías buscar patrones específicos en los datos para determinar el tipo.
	return true
}

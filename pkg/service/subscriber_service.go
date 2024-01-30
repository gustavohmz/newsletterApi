// pkg/service/subscriber_service.go

package service

import (
	"newsletter-app/pkg/domain"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"time"
)

// SubscriberService es una estructura que maneja la lógica de negocio relacionada con los suscriptores y boletines.
type SubscriberService struct {
	subscriberRepository *mongodb.SubscriberRepository
}

// NewSubscriberService crea una nueva instancia de SubscriberService.
func NewSubscriberService() *SubscriberService {
	return &SubscriberService{
		subscriberRepository: mongodb.NewSubscriberRepository(),
	}
}

// Subscribe agrega un nuevo suscriptor.
func (s *SubscriberService) Subscribe(email string, category string) error {
	subscriber := domain.Subscriber{
		Email:            email,
		SubscriptionDate: time.Now(),
		Category:         category,
	}

	return s.subscriberRepository.SaveSubscriber(subscriber)
}

// Unsubscribe elimina un suscriptor por su dirección de correo electrónico.
func (s *SubscriberService) Unsubscribe(email string) error {
	// Lógica para manejar la cancelación de la suscripción
	return s.subscriberRepository.DeleteSubscriberByEmail(email)
}

// GetSubscriberByEmail obtiene un suscriptor por dirección de correo electrónico.
func (s *SubscriberService) GetSubscriberByEmail(email string) (*domain.Subscriber, error) {
	return s.subscriberRepository.GetSubscriberByEmail(email)
}

// GetSubscribers obtiene la lista de suscriptores.
func (s *SubscriberService) GetSubscribers() ([]domain.Subscriber, error) {
	return s.subscriberRepository.GetSubscribers()
}

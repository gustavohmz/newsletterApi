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

// Unsubscribe elimina un suscriptor por su dirección de correo electrónico y categoria.
func (s *SubscriberService) Unsubscribe(email, category string) error {
	return s.subscriberRepository.DeleteSubscriberByEmail(email, category)
}

// GetSubscriberByEmail obtiene un suscriptor por dirección de correo electrónico y categoria.
func (s *SubscriberService) GetSubscriberByEmail(email, category string) (*domain.Subscriber, error) {
	return s.subscriberRepository.GetSubscriberByEmailAndCategory(email, category)
}

// GetSubscribers obtiene la lista de suscriptores con parámetros de búsqueda y paginación.
func (s *SubscriberService) GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error) {
	return s.subscriberRepository.GetSubscribers(email, category, page, pageSize)
}

package mongodb

import (
	"context"
	"newsletter-app/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SubscriberRepository es una estructura que maneja operaciones relacionadas con los suscriptores y boletines en MongoDB.
type SubscriberRepository struct {
	subscriberCollection *mongo.Collection
}

// NewSubscriberRepository crea una nueva instancia de SubscriberRepository.
func NewSubscriberRepository() *SubscriberRepository {
	return &SubscriberRepository{
		subscriberCollection: client.Database("newsletter-app").Collection("subscribers"),
	}
}

// SaveSubscriber guarda un nuevo suscriptor en la base de datos.
func (r *SubscriberRepository) SaveSubscriber(subscriber domain.Subscriber) error {
	_, err := r.subscriberCollection.InsertOne(context.TODO(), subscriber)
	return err
}

// GetSubscriberByEmail obtiene un suscriptor por dirección de correo electrónico.
func (r *SubscriberRepository) GetSubscriberByEmail(email string) (*domain.Subscriber, error) {
	var subscriber domain.Subscriber
	err := r.subscriberCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&subscriber)
	if err != nil {
		return nil, err
	}
	return &subscriber, nil
}

// GetSubscribers obtiene la lista de suscriptores desde la base de datos.
func (r *SubscriberRepository) GetSubscribers() ([]domain.Subscriber, error) {
	var subscribers []domain.Subscriber

	cursor, err := r.subscriberCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var subscriber domain.Subscriber
		if err := cursor.Decode(&subscriber); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}

	return subscribers, nil
}

// DeleteSubscriberByEmail elimina un suscriptor por su dirección de correo electrónico.
func (r *SubscriberRepository) DeleteSubscriberByEmail(email string) error {
	_, err := r.subscriberCollection.DeleteOne(context.TODO(), bson.M{"email": email})
	return err
}

package mongodb

import (
	"context"
	"newsletter-app/pkg/domain"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SubscriberRepository es una estructura que maneja operaciones relacionadas con los suscriptores y boletines en MongoDB.
type SubscriberRepository struct {
	subscriberCollection *mongo.Collection
}

// NewSubscriberRepository crea una nueva instancia de SubscriberRepository.
func NewSubscriberRepository() *SubscriberRepository {
	mongoDb := os.Getenv("mongoDb")
	mongoSubscriberCollection := os.Getenv("mongoSubscriberCollection")

	return &SubscriberRepository{
		subscriberCollection: client.Database(mongoDb).Collection(mongoSubscriberCollection),
	}
}

// SaveSubscriber guarda un nuevo suscriptor en la base de datos.
func (r *SubscriberRepository) SaveSubscriber(subscriber domain.Subscriber) error {
	_, err := r.subscriberCollection.InsertOne(context.TODO(), subscriber)
	return err
}

// GetSubscriberByEmailAndCategory obtiene un suscriptor por su dirección de correo electrónico y categoría.
func (r *SubscriberRepository) GetSubscriberByEmailAndCategory(email, category string) (*domain.Subscriber, error) {
	var subscriber domain.Subscriber
	filter := bson.M{"email": email, "category": category}
	err := r.subscriberCollection.FindOne(context.TODO(), filter).Decode(&subscriber)
	if err != nil {
		return nil, err
	}
	return &subscriber, nil
}

// GetSubscribers obtiene la lista de suscriptores con parámetros de búsqueda y paginación.
func (r *SubscriberRepository) GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error) {
	var subscribers []domain.Subscriber

	filter := bson.M{}
	if email != "" {
		filter["email"] = email
	}
	if category != "" {
		filter["category"] = category
	}

	// Puedes agregar lógica adicional para la paginación aquí si es necesario

	cursor, err := r.subscriberCollection.Find(context.TODO(), filter)
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

// DeleteSubscriberByEmailAndCategory elimina un suscriptor por su dirección de correo electrónico y/o categoría.
func (r *SubscriberRepository) DeleteSubscriberByEmail(email, category string) error {
	filter := bson.M{"email": email}

	// Agregar la condición de categoría si está presente
	if category != "" {
		filter["category"] = category
	}

	_, err := r.subscriberCollection.DeleteMany(context.TODO(), filter)
	return err
}

// GetSubscribersByCategory obtiene suscriptores por categoría.
func (r *SubscriberRepository) GetSubscribersByCategory(category string) ([]domain.Subscriber, error) {
	filter := bson.M{"category": category}

	cursor, err := r.subscriberCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var subscribers []domain.Subscriber
	for cursor.Next(context.TODO()) {
		var subscriber domain.Subscriber
		if err := cursor.Decode(&subscriber); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}

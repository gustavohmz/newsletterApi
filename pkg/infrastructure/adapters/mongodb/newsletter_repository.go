package mongodb

import (
	"context"
	"newsletter-app/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewsletterRepository es una estructura que maneja operaciones relacionadas con los boletines en MongoDB.
type NewsletterRepository struct {
	newsletterCollection *mongo.Collection
}

// NewNewsletterRepository crea una nueva instancia de NewsletterRepository.
func NewNewsletterRepository() *NewsletterRepository {
	return &NewsletterRepository{
		newsletterCollection: client.Database("newsletter-app").Collection("newsletters"),
	}
}

// SaveNewsletter guarda un nuevo boletín en la base de datos.
func (r *NewsletterRepository) SaveNewsletter(newsletter domain.Newsletter) error {
	_, err := r.newsletterCollection.InsertOne(context.TODO(), newsletter)
	return err
}

// GetNewsletterByID obtiene un boletín por ID.
func (r *NewsletterRepository) GetNewsletterByID(newsletterID string) (*domain.Newsletter, error) {
	var newsletter domain.Newsletter
	objectID, err := primitive.ObjectIDFromHex(newsletterID)
	if err != nil {
		return nil, err
	}

	err = r.newsletterCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&newsletter)
	if err != nil {
		return nil, err
	}
	return &newsletter, nil
}

// UpdateNewsletter actualiza un boletín existente en la base de datos.
func (r *NewsletterRepository) UpdateNewsletter(newsletterID string) error {
	// Crea un filtro para identificar el boletín por su ID
	objectID, err := primitive.ObjectIDFromHex(newsletterID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}

	// Crea una actualización para establecer el campo "sent" en true
	update := bson.M{"$set": bson.M{"sent": true}}

	// Ejecuta la actualización en la base de datos
	_, err = r.newsletterCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

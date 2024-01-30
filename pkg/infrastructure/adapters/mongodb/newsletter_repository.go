package mongodb

import (
	"context"
	"newsletter-app/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (r *NewsletterRepository) UpdateNewsletter(newsletter domain.Newsletter) error {
	// Crea un filtro para identificar el boletín por su ID
	filter := bson.M{"_id": newsletter.ID}

	// Crea una actualización para establecer los campos proporcionados
	update := bson.M{"$set": bson.M{
		"name":        newsletter.Name,
		"category":    newsletter.Category,
		"subject":     newsletter.Subject,
		"content":     newsletter.Content,
		"attachments": newsletter.Attachments,
	}}

	// Ejecuta la actualización en la base de datos
	_, err := r.newsletterCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

// GetNewsletterByCategory obtiene un boletín por categoría.
func (r *NewsletterRepository) GetNewsletterByCategory(category string) (*domain.Newsletter, error) {
	var newsletter domain.Newsletter
	filter := bson.M{"category": category}

	err := r.newsletterCollection.FindOne(context.TODO(), filter).Decode(&newsletter)
	if err != nil {
		// Si no se encuentra un boletín con la categoría dada, retornamos nil y sin error
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &newsletter, nil
}

// GetNewsletters obtiene una lista de boletines con opciones de búsqueda y paginación.
func (r *NewsletterRepository) GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error) {
	// Crea un filtro para la búsqueda por nombre (si se proporciona)
	filter := bson.M{}
	if searchName != "" {
		filter["name"] = primitive.Regex{Pattern: searchName, Options: "i"}
	}

	// Opciones de búsqueda
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	// Realiza la búsqueda en la base de datos
	cursor, err := r.newsletterCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Declara una variable para almacenar los boletines
	var newsletters []domain.Newsletter

	// Itera sobre los resultados y decodifica cada documento en una estructura de boletín
	for cursor.Next(context.TODO()) {
		var newsletter domain.Newsletter
		if err := cursor.Decode(&newsletter); err != nil {
			return nil, err
		}
		newsletters = append(newsletters, newsletter)
	}

	// Maneja cualquier error que pueda ocurrir durante el cursor
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return newsletters, nil
}

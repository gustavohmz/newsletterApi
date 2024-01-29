package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Subscriber representa un suscriptor del boletín.
// swagger:model
type Subscriber struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email            string             `json:"email"`
	SubscriptionDate time.Time          `json:"subscription_date"`
}

type Subscribers []Subscriber

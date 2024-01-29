package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category representa las categorías posibles para un boletín.
type Category int

const (
	SpecialOffers     Category = 1
	Memberships       Category = 2
	MonthlyPromotions Category = 3
	NewProducts       Category = 4
)

// Newsletter representa un boletín.
// swagger:model
type Newsletter struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:""`
	Name        string             `json:"name"`
	Categories  []Category         `json:"categories"`
	Attachments []Attachment       `json:"attachments"`
}

// Attachment representa un archivo adjunto en el boletín.
// swagger:model
type Attachment struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// SentRecipient almacena información sobre a quién se envió el boletín.
type SentRecipient struct {
	Email string `json:"email"`
}

package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Newsletter representa un boletín.
// swagger:model
type Newsletter struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:""`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Content     string             `json:"content"`
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

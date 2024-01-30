package request

import "go.mongodb.org/mongo-driver/bson/primitive"

// UpdateNewsletterRequest representa la estructura para la solicitud de actualización del boletín.
type UpdateNewsletterRequest struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Subject     string             `json:"subject"`
	Content     string             `json:"content"`
	Attachments []Attachment       `json:"attachments"`
}

// Attachment representa un archivo adjunto en el boletín.
// swagger:model
type Attachment struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Type string `json:"type"`
}

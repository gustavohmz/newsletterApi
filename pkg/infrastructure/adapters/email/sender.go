package email

import (
	"crypto/tls"
	"encoding/base64"
	"io"
	"newsletter-app/pkg/domain"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Sender define una interfaz para enviar correos electrónicos
type Sender interface {
	Send(subject, body string, to []string, attachments []*domain.Attachment) error
}

// Implementación concreta utilizando el servicio de Brevo
type BrevoEmailSender struct {
}

// NewBrevoEmailSender crea una nueva instancia de BrevoEmailSender
func NewBrevoEmailSender() *BrevoEmailSender {
	return &BrevoEmailSender{}
}

// Send implementa la interfaz Sender
func (b *BrevoEmailSender) Send(subject, body string, to []string, attachments []*domain.Attachment) error {
	emailSender := "gustavohdzmz@gmail.com"
	emailPass := "ZsM7SmW9NDX063Yd"
	smtpServer := "smtp-relay.brevo.com"
	smtpPort := "587"

	smtpPortInt, err := strconv.Atoi(smtpPort)
	if err != nil {
		return err
	}

	d := gomail.NewDialer(smtpServer, smtpPortInt, emailSender, emailPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailSender)
	mailer.SetHeader("To", to...)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	// Adjuntar archivos
	for _, attachment := range attachments {
		// Convertir la cadena Base64 a datos binarios
		data, err := base64.StdEncoding.DecodeString(attachment.Data)
		if err != nil {
			return err
		}

		// Adjuntar el archivo al mensaje
		mailer.Attach(attachment.Name, gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(data)
			return err
		}))
	}

	if err := d.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

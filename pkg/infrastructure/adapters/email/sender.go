package email

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Sender define una interfaz para enviar correos electrónicos
type Sender interface {
	Send(subject, body string, to []string) error
}

// Implementación concreta utilizando el servicio de Brevo o cualquier otro proveedor de correo
type BrevoEmailSender struct {
	// Puedes agregar configuraciones específicas del proveedor aquí
}

// NewBrevoEmailSender crea una nueva instancia de BrevoEmailSender
func NewBrevoEmailSender() *BrevoEmailSender {
	return &BrevoEmailSender{}
}

// Send implementa la interfaz Sender
func (b *BrevoEmailSender) Send(subject, body string, to []string) error {
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

	if err := d.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

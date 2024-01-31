# Usa la imagen oficial de Go 1.21.6
FROM golang:1.21.6

# Establece el directorio de trabajo en /app
WORKDIR /app

# descarga  los modulos
COPY go.mod go.sum ./
RUN go mod download

COPY . .
EXPOSE 8080
# Compila la aplicación
RUN go build -o newsletter
ENV \
    mongoUrl="mongodb+srv://gustavohdzmz:COERlJXgVI3XSp6M@newsletter.9soh00l.mongodb.net/?retryWrites=true&w=majority" \
    mongoDb="newsletter-app" \
    mongoNewsletterCollection="newsletters" \
    mongoSubscriberCollection="subscribers" \
    emailSender="gustavotest025@gmail.com" \
    emailPass="UAzmCB4Wba5VTxI7" \
    smtpServer="smtp-relay.brevo.com" \
    smtpPort=587

# Ejecuta la aplicación al iniciar el contenedor
CMD ["./newsletter"]
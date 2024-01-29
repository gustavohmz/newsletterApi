// pkg/infrastructure/adapters/mongodb/mongodb.go

package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    context.Context
)

// Connect establece una conexión a MongoDB
func Connect(connectionString string) error {
	// Crear el cliente de MongoDB
	clientOptions := options.Client().ApplyURI(connectionString)
	newClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar la conexión
	err = newClient.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("Error al hacer ping a MongoDB: %v", err)
	}

	client = newClient
	return nil
}

// GetClient devuelve el cliente MongoDB
func GetClient() *mongo.Client {
	return client
}

// Disconnect cierra la conexión a MongoDB
func Disconnect() error {
	if client != nil {
		return client.Disconnect(ctx)
	}
	return nil
}

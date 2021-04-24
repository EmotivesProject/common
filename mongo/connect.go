package mongo

import (
	"context"
	"log"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI    string
	DBName string
}

var (
	db *mongo.Database
)

func Connect(config Config) {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.URI)

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Connected to MongoDB!")

	db = client.Database(config.DBName)
}

func GetDatabase() *mongo.Database {
	return db
}

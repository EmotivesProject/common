package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	URI    string
	DBName string
}

const (
	retries   = 50
	sleepTime = 5
)

var (
	dbConfig           Config
	db                 *mongo.Database
	errFailedToConnect = errors.New("Failed to connect to db")
)

func Connect(config Config) error {
	dbConfig = config

	var client *mongo.Client

	var err error

	// Set client options
	clientOptions := options.Client().ApplyURI(config.URI)

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	for i := 0; i < retries; i++ {
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			time.Sleep(sleepTime * time.Second)

			continue
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			time.Sleep(sleepTime * time.Second)

			continue
		}

		// If it gets here no errors has happened
		break
	}

	if client == nil {
		return errFailedToConnect
	}

	logger.Info("Connected to MongoDB!")

	db = client.Database(config.DBName)

	return nil
}

func GetDatabase() *mongo.Database {
	if db == nil {
		if err := Connect(dbConfig); err != nil {
			logger.Error(errFailedToConnect)
		}
	}

	return db
}

func CloseDatabase() *mongo.Database {
	if db == nil {
		_ = db.Client().Disconnect(context.Background())
	}

	return db
}

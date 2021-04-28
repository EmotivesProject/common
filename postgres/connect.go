package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/jackc/pgx/v4"
)

type Config struct {
	URI string
}

const (
	retries   = 50
	sleepTime = 5
)

var (
	dbConfig           Config
	db                 *pgx.Conn
	errFailedToConnect = errors.New("Failed to connect to db")
)

func Connect(config Config) error {
	dbConfig = config

	var conn *pgx.Conn

	var err error

	for i := 0; i < retries; i++ {
		conn, err = pgx.Connect(context.Background(), config.URI)
		if err != nil {
			logger.Error(err)
			time.Sleep(sleepTime * time.Second)

			continue
		}

		break
	}

	if err != nil {
		return errFailedToConnect
	}

	logger.Info("Successfully connected to the database")

	db = conn

	return nil
}

func GetDatabase() *pgx.Conn {
	if db == nil {
		if err := Connect(dbConfig); err != nil {
			logger.Error(errFailedToConnect)
		}
	}

	return db
}

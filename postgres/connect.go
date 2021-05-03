package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/jackc/pgx/v4/pgxpool"
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
	db                 *pgxpool.Pool
	errFailedToConnect = errors.New("Failed to connect to db")
)

func Connect(config Config) error {
	dbConfig = config

	var pool *pgxpool.Pool

	var err error

	for i := 0; i < retries; i++ {
		pool, err = pgxpool.Connect(context.Background(), config.URI)
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

	db = pool

	return nil
}

func GetDatabase() *pgxpool.Pool {
	if db == nil {
		if err := Connect(dbConfig); err != nil {
			logger.Error(errFailedToConnect)
		}
	}

	return db
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}

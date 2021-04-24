package postgres

import (
	"context"

	"github.com/TomBowyerResearchProject/common/logger"
	"github.com/jackc/pgx/v4"
)

type Config struct {
	URI string
}

var (
	db *pgx.Conn
)

func Connect(config Config) {
	conn, err := pgx.Connect(context.Background(), config.URI)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Successfully connected to the database")
	db = conn
}

func GetDatabase() *pgx.Conn {
	return db
}

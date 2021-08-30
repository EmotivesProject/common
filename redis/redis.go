package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr   string
	Prefix string
}

var (
	db                 *redis.Client
	prefix             string
	errFailedToConnect = errors.New("Failed to connect to db")
)

func Init(config Config) error {
	if db != nil {
		return nil
	}

	prefix = config.Prefix

	// If not set create a new client
	db = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if db == nil {
		return errFailedToConnect
	}

	return nil
}

func Set(ctx context.Context, key string, value interface{}, timeout time.Duration) error {
	return SetEx(ctx, key, value, 0)
}

func SetEx(ctx context.Context, key string, value interface{}, timeout time.Duration) error {
	fullKey := fmt.Sprintf("%s.%s", prefix, key)

	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return db.Set(ctx, fullKey, data, timeout).Err()
}

func Get(ctx context.Context, key string, model interface{}) (interface{}, error) {
	fullKey := fmt.Sprintf("%s.%s", prefix, key)

	result, err := db.Get(ctx, fullKey).Result()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(result), &model)

	return model, err
}

func Down() {
	if db != nil {
		db.Close()
	}
}

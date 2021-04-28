package kafka

import (
	"time"

	"github.com/TomBowyerResearchProject/common/logger"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Handler func([]byte)

type ConfigConsumer struct {
	Topic  string
	Server string
	Group  string
	Handle Handler
}

var (
	kafkaConsumer  *kafka.Consumer
	consumerConfig ConfigConsumer
)

func InitConsumer(configConsumer ConfigConsumer) error {
	var consumer *kafka.Consumer

	var err error

	consumerConfig = configConsumer

	for i := 0; i < retries; i++ {
		consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": consumerConfig.Server,
			"group.id":          consumerConfig.Group,
		})
		if err == nil {
			break
		}

		logger.Error(err)
		time.Sleep(sleepTime * time.Second)
	}

	if err != nil {
		return err
	}

	logger.Info("Connected to kafka")

	kafkaConsumer = consumer

	return nil
}

func Run() error {
	if kafkaConsumer == nil {
		if err := InitConsumer(consumerConfig); err != nil {
			return err
		}
	}

	err := kafkaConsumer.Subscribe(consumerConfig.Topic, nil)
	if err != nil {
		logger.Error(err)
	}

	defer kafkaConsumer.Close()

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			consumerConfig.Handle(msg.Value)
		} else {
			logger.Error(err)
		}
	}
}

func CloseConsumer() {
	if kafkaConsumer != nil {
		kafkaConsumer.Close()
	}
}

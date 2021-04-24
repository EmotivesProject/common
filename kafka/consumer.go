package kafka

import (
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

func InitConsumer(configConsumer ConfigConsumer) {
	consumerConfig = configConsumer

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": consumerConfig.Server,
		"group.id":          consumerConfig.Group,
	})
	if err != nil {
		logger.Error(err)
	}

	kafkaConsumer = consumer
}

func Run() {
	err := kafkaConsumer.Subscribe(consumerConfig.Topic, nil)
	if err != nil {
		logger.Error(err)
	}

	logger.Info("Connected to kafka")
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

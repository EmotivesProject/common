package kafka

import (
	"encoding/json"

	"github.com/TomBowyerResearchProject/common/logger"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type ConfigProducer struct {
	Topic  string
	Server string
}

var (
	kafkaProducer  *kafka.Producer
	producerConfig ConfigProducer
)

func InitProducer(configProducer ConfigProducer) {
	producerConfig = configProducer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": producerConfig.Server,
	})
	if err != nil {
		logger.Error(err)
	}

	kafkaProducer = producer
}

func ProduceEvent(event interface{}) {
	stringEvent, err := json.Marshal(event)
	if err != nil {
		logger.Error(err)
	}

	err = kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &producerConfig.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(stringEvent)},
		nil,
	)

	if err != nil {
		logger.Error(err)
	} else {
		logger.Infof("Sent event off to kafka %s", event)
	}
}

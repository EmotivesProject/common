package kafka

import (
	"encoding/json"
	"time"

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

func InitProducer(configProducer ConfigProducer) error {
	var producer *kafka.Producer

	var err error

	producerConfig = configProducer

	for i := 0; i < retries; i++ {
		producer, err = kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": producerConfig.Server,
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

	kafkaProducer = producer

	return nil
}

func ProduceEvent(event interface{}) error {
	if kafkaProducer == nil {
		if err := InitProducer(producerConfig); err != nil {
			return err
		}
	}

	stringEvent, err := json.Marshal(event)
	if err != nil {
		logger.Error(err)
	}

	err = kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &producerConfig.Topic, Partition: kafka.PartitionAny},
		Value:          stringEvent,
	},
		nil,
	)

	if err != nil {
		logger.Error(err)

		return err
	}

	logger.Infof("Sent event off to kafka %s", event)

	return nil
}

func CloseProducer() {
	if kafkaProducer != nil {
		kafkaProducer.Close()
	}
}

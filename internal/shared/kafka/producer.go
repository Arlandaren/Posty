package kafka

import (
	"github.com/IBM/sarama"
	"log"
)

type Producer struct {
	config       *Config
	syncProducer sarama.SyncProducer
}

func NewProducer(config *Config) (*Producer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(config.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	return &Producer{
		config:       config,
		syncProducer: producer,
	}, nil
}

func (p *Producer) ProduceMessage(topic string, message []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := p.syncProducer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to produce message: %v", err)
	}

	return err
}

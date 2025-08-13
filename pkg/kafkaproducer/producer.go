package kafkaproducer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	Producer *kafka.Producer
}

type KafkaProducerConfig struct {
	Servers  string
	ClientId string
	Acks     string
}

var DefaultConfig = KafkaProducerConfig{
	Servers:  "localhost:9092",
	ClientId: "alpaca-consumer",
	Acks:     "all",
}

func NewKafkaProducer(config KafkaProducerConfig) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Servers,
		"client.id":        config.ClientId,
		"acks":             config.Acks,
	})
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{Producer: producer}, nil
}

func (p *KafkaProducer) Produce(topic string, key []byte, value []byte) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value,
	}, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}
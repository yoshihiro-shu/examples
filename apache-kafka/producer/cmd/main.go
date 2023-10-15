package main

import (
	"log"

	"github.com/IBM/sarama"
)

type Producer struct {
	sarama.SyncProducer
}

func NewProducer() (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
		return nil, err
	}

	return &Producer{producer}, nil
}

func (p *Producer) SendMessages(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}
	_, _, err := p.SyncProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	producer, err := NewProducer()
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	log.Printf("==> start producing %#v\n", producer)
	log.Printf("==> status of producing %#v", producer.TxnStatus())

	topic := "news"
	key := []byte("key")
	value := []byte("value")
	err = producer.SendMessages(topic, key, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("success")
}

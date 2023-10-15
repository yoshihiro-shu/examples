package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

// this is worked. but consume every messages
// func main() {
// 	brokers := []string{"localhost:9092"} // Kafkaブローカーのアドレス
// 	topic := "news"                       // ここにトピック名を指定

// 	config := sarama.NewConfig()
// 	config.Consumer.Return.Errors = true

// 	// Consumerの作成
// 	consumer, err := sarama.NewConsumer(brokers, config)
// 	if err != nil {
// 		log.Fatalf("Error creating consumer: %s", err)
// 	}
// 	defer consumer.Close()

// 	// トピックのパーティションを取得
// 	partitions, err := consumer.Partitions(topic)
// 	if err != nil {
// 		log.Fatalf("Error retrieving partitions: %s", err)
// 	}

// 	// 各パーティションからメッセージを消費
// 	for _, partition := range partitions {
// 		pc, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
// 		if err != nil {
// 			log.Fatalf("Error consuming partition %d: %s", partition, err)
// 		}
// 		defer pc.Close()

// 		go func(pc sarama.PartitionConsumer) {
// 			for message := range pc.Messages() {
// 				fmt.Printf("Received message from partition %d at offset %d: %s\n", message.Partition, message.Offset, string(message.Value))
// 			}
// 		}(pc)
// 	}

// 	// シグナルを待機（Ctrl+Cで終了）
// 	sigchan := make(chan os.Signal, 1)
// 	signal.Notify(sigchan, os.Interrupt)
// 	<-sigchan
// }

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	brokers := []string{"localhost:9092"}
	group := "example-group"
	topic := "news"

	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // Kafkaのバージョンに合わせて変更
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}

	ctx := context.Background()
	handler := exampleConsumerGroupHandler{}

	// シグナルを待機（Ctrl+Cで終了）
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Consumer Groupでのメッセージの消費を開始
	go func() {
		for {
			err := client.Consume(ctx, []string{topic}, handler)
			if err != nil {
				log.Fatalf("Error from consumer: %v", err)
			}
		}
	}()

	<-sigchan
	client.Close()
}

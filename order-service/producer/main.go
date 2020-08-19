package main

import (
	"fmt"
	"time"

	"log"

	"microservice-app/order-service/producer/service"

	"github.com/Shopify/sarama"
)

func main() {
	kafkaConfig := getKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		log.Printf("Unable too create kafka producer got error %v", err)
		return
	}
	defer func() {
		if err := producers.Close(); err != nil {
			log.Printf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	log.Println("Success create kafka sync-producer")

	kafka := &service.KafkaProducer{
		Producer: producers,
	}

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("message number %v", i)
		err := kafka.SendMessage("test_topic", msg)
		if err != nil {
			panic(err)
		}
	}
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

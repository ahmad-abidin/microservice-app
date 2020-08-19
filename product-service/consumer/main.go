package main

import (
	"os"
	"time"

	"log"
	"microservice-app/product-service/consumer/service"

	"github.com/Shopify/sarama"
)

func main() {
	kafkaConfig := getKafkaConfig("", "")

	consumers, err := sarama.NewConsumer([]string{"kafka:9092"}, kafkaConfig)
	if err != nil {
		log.Printf("Error create kakfa consumer got error %v", err)
	}
	defer func() {
		if err := consumers.Close(); err != nil {
			log.Fatalln(err)
			return
		}
	}()

	kafkaConsumer := &service.KafkaConsumer{
		Consumer: consumers,
	}

	signals := make(chan os.Signal, 1)
	kafkaConsumer.Consume([]string{"test_topic"}, signals)
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

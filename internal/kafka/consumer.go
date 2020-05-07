package kafka

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

func Init() (string, *kafka.Consumer){
	broker := os.Getenv("BOOTSTRAP_SERVER")
	topicJobOffer := "com.github.JobOffer"
	topics := []string{topicJobOffer}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     broker,
		"broker.address.family": "v4",
		"group.id":              "Recruitment",
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest",
	})
	if err != nil {
		log.Fatal("Failed to create consumer", err)
	}
	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatal("Failed to subscribe to topics", err)
	}
	return topicJobOffer, c
}
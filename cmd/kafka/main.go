package main

import (
	"MyFirstGoWebServer/internal/core"
	"MyFirstGoWebServer/internal/database"
	_jobOfferConsumer "MyFirstGoWebServer/internal/joboffer/consumer"
	_jobOfferRepo "MyFirstGoWebServer/internal/joboffer/repository"
	_jobOfferService "MyFirstGoWebServer/internal/joboffer/service"
	_kafkaConsumer "MyFirstGoWebServer/internal/kafka"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"os"
)

func main() {
	log.Println("Startup kafka main")
	db := database.Init()

	topicJobOffer, c := _kafkaConsumer.Init()
	defer c.Close()

	jobOfferRepo := _jobOfferRepo.NewJobOfferRepository(db)
	jobOfferService := _jobOfferService.NewJobOfferService(jobOfferRepo)
	jobOfferConsumer := _jobOfferConsumer.NewJobOfferConsumer(jobOfferService)

	sigChan := make(chan os.Signal, 1)
	run := true
	log.Println("Start to poll messages...")
	for run == true {
		select {
		case sig := <-sigChan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			event := c.Poll(100)
			if event == nil {
				continue
			}
			HandleMessageFactory(event, topicJobOffer, jobOfferConsumer)
		}
	}
}

func HandleMessageFactory(event kafka.Event, topicJobOffer string, jobOfferConsumer core.JobOfferConsumer) {
	switch e := event.(type) {
	case *kafka.Message:
		log.Println("topic", *e.TopicPartition.Topic)
		switch *e.TopicPartition.Topic {
		case topicJobOffer:
			jobOfferConsumer.HandleMessage(e.Key, e.Value)
		}
	}
}

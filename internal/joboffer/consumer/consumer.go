package consumer

import (
	"MyFirstGoWebServer/internal/core"
	"context"
	"encoding/json"
	"log"
)

type jobOfferConsumerImpl struct {
	JobOfferService core.JobOfferService
}

func NewJobOfferConsumer(js core.JobOfferService) core.JobOfferConsumer {
	return &jobOfferConsumerImpl{
		JobOfferService: js,
	}
}

func (h *jobOfferConsumerImpl) HandleMessage(key []byte, value []byte) {
	var jobOffer core.JobOfferRepresentation
	type JSON map[string]string
	var generic JSON
	if err := json.Unmarshal(value, &jobOffer); err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(key, &generic); err != nil {
		log.Println(err)
		return
	}
	if id, ok := generic["id"]; ok {
		_ = h.JobOfferService.SaveKafkaMessage(context.Background(), id, &jobOffer)
	}
}

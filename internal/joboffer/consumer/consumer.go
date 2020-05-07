package consumer

import (
	"MyFirstGoWebServer/internal/core"
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
	k := string(key)
	v := string(value)
	log.Println("message: ", k, v)
}

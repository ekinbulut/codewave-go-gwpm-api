package main

import (
	"hermes/cmd/config"
	"hermes/contracts/data"
	"hermes/internals/services"
	"log"
)

type Tracker struct {
	config *config.ConsumerConfigurations
}

func NewTracker(config *config.ConsumerConfigurations) *Tracker {
	return &Tracker{
		config: config,
	}
}

func (t *Tracker) Consume() {

	message := make(chan data.DataContract)

	consumer := services.NewConsumer(&t.config.Rabbitmq)
	go consumer.Consume(message)

	<-message

	// log print message
	log.Printf("Consumer message: %+v", message)

}

package services

import (
	"context"
	"encoding/json"
	"fmt"
	"hermes/cmd/config"
	"hermes/contracts/data"
	"hermes/internals/api/presenter"

	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IPublisher interface {
	Publish(message presenter.Message) (uuid.UUID, error)
}

type Publisher struct {
	connection *amqp.Connection
	config     *config.RabbitMqConfiguration
}

func NewPublisher(config *config.RabbitMqConfiguration) IPublisher {

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil
	}
	return &Publisher{
		connection: conn,
		config:     config,
	}
}

func (r *Publisher) Publish(message presenter.Message) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ch, err := r.connection.Channel()
	if err != nil {
		return uuid.Nil, err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(r.config.Queue, true, false, false, false, nil)
	if err != nil {
		return uuid.Nil, err
	}

	// create a data contract
	dataContract := data.NewDataContract(message)

	pmessage, _ := json.MarshalIndent(dataContract, "", " ")

	ch.PublishWithContext(ctx, "", r.config.Queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        pmessage,
	})

	return dataContract.TransactionId, nil
}

func (r *Publisher) Close() {
	r.connection.Close()
}

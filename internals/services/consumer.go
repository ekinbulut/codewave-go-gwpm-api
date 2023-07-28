package services

import (
	"encoding/json"
	"fmt"
	"hermes/cmd/config"
	"hermes/contracts/data"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IConsumer interface {
	Consume(messageChan chan data.DataContract)
}

type Consumer struct {
	connection *amqp.Connection
	config     *config.RabbitMqConfiguration
}

func NewConsumer(config *config.RabbitMqConfiguration) IConsumer {
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.Username, config.Password, config.Host, config.Port)
	log.Printf("Connecting to RabbitMQ: %s", connectionString)
	var err error
	if connection == nil {
		connection, err = amqp.Dial(connectionString)
		if err != nil {
			failOnError(err, "Failed to connect to RabbitMQ")
			defer connection.Close()
			return nil
		}
	}
	return &Consumer{
		connection: connection,
		config:     config,
	}
}

func (c *Consumer) Consume(messageChan chan data.DataContract) {
	ch, err := c.connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer c.connection.Close()

	q, err := ch.QueueDeclare(
		c.config.Queue, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var dataContract data.DataContract
			err := json.Unmarshal(d.Body, &dataContract)
			failOnError(err, "Failed to unmarshal")
			time.Sleep(2 * time.Second)
			log.Printf("Done")
			d.Ack(true)
			messageChan <- dataContract
		}
	}()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

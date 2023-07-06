package services

import (
	"hermes/cmd/config"
	"hermes/internals/api/presenter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayPublishMessage(t *testing.T) {

	var config = &config.RabbitMqConfiguration{
		Port:     8080,
		Host:     "localhost",
		Username: "guest",
		Queue:    "test",
	}

	message := &presenter.Message{
		Numbers: []string{"11111111111"},
		Template: struct {
			Name string `json:"name"`
		}{
			Name: "test",
		},
	}
	relayer := NewPublisher(config)
	_, err := relayer.Publish(*message)

	assert.NoError(t, err, "Publish should not fail")
}

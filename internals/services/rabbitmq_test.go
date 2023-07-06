package services

import (
	"hermes/internals/api/presenter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayCreate(t *testing.T) {

	var config = &Config{
		Port:      8080,
		Host:      "localhost",
		Username:  "guest",
		QueueName: "test",
	}

	relayer := NewRelayer(config)

	assert.NotEqual(t, relayer.connection, nil, "Connection should not be nil")
}

func TestRelayPublishMessage(t *testing.T) {

	var config = &Config{
		Port:      8080,
		Host:      "localhost",
		Username:  "guest",
		QueueName: "test",
	}

	message := &presenter.Message{
		Message: "test",
	}
	relayer := NewRelayer(config)
	err := relayer.Publish(*message)

	assert.NoError(t, err, "Publish should not fail")
}

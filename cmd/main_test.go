package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	_, err := ReadConfigs()

	

	assert.Error(t, err)

	//fmt.Printf("Configuration: %+v", configuration)
	//assert.Equal(t, 8080, configuration.Server.Port)
}

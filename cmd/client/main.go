package main

import (
	"fmt"
	c "hermes/cmd/config"
	"log"

	"github.com/spf13/viper"
)

func main() {
	var config, err = ReadConfigs()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
	client := NewTracker(&config)
	client.Consume()
}

func ReadConfigs() (config c.ConsumerConfigurations, err error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.ConsumerConfigurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return configuration, err
	}

	log.Printf("%v", configuration.Rabbitmq.Host)

	return configuration, nil
}

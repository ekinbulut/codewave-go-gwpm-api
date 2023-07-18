package main

import (
	"fmt"

	c "hermes/cmd/config"
	"hermes/cmd/server"

	"github.com/spf13/viper"
)

func main() {

	configuration, err := ReadConfigs()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	app := server.NewServer(&configuration)
	app.Run()
}

func ReadConfigs() (config c.Configurations, err error) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return configuration, err
	}

	return configuration, nil
}

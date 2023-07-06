package config

type Configurations struct {
	Server   ServerConfiguration
	Rabbitmq RabbitMqConfiguration
}

type RabbitMqConfiguration struct {
	Port     int
	Host     string
	Username string
	Password string
	Queue    string
	Exchange string
}

type ServerConfiguration struct {
	Port int
	Host string
	Env  string
}

package config

type Configurations struct {
	Server    ServerConfiguration
	Rabbitmq  RabbitMqConfiguration
	RateLimit RateLimiterConfigurations
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

type RateLimiterConfigurations struct {
	Max        int
	Expiration int
}

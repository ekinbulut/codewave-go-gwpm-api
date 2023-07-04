package main

type Config struct {
	Port int
	Host string
}

func NewConfig() *Config {
	return &Config{
		Port: 8080,
		Host: "localhost",
	}
}

// read configs
func (c *Config) ReadConfig() {
	// read a config file from path
}

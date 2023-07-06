package server

import (
	"fmt"
	"hermes/cmd/config"
	"hermes/internals/api/routes"
	"hermes/internals/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config *config.Configurations
	app    *fiber.App
}

func NewServer(config *config.Configurations) *Server {
	return &Server{
		config: config,
		app:    fiber.New(),
	}
}

func (s *Server) Run() {

	if s.config == nil {
		s.config = &config.Configurations{
			Server: config.ServerConfiguration{
				Port: 8080,
				Host: "localhost",
			},
		}
	}

	if s.config.Server.Port == 0 {
		s.config.Server.Port = 8080
	}

	api := s.app.Group("/api/v1")

	publisher := services.NewRelayer(&s.config.Rabbitmq)

	// register routes
	routes.MessageRouter(api, publisher)

	port := s.config.Server.Port
	fmt.Printf("Server running on port %d\n", port)

	log.Fatal(s.app.Listen(fmt.Sprintf(":%d", port)))
}

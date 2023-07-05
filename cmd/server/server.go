package server

import (
	"fmt"
	"hermes/internals/api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config *Config
	app    *fiber.App
}

func NewServer(config *Config) *Server {
	return &Server{
		config: config,
		app:    fiber.New(),
	}
}

func (s *Server) Run() {

	if s.config == nil {
		s.config = NewConfig()
	}

	if s.config.Port == 0 {
		s.config.Port = 8080
	}

	api := s.app.Group("/api/v1")

	// register routes
	routes.MessageRouter(api)

	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	port := s.config.Port
	fmt.Printf("Server running on port %d\n", port)

	log.Fatal(s.app.Listen(fmt.Sprintf(":%d", port)))
}

package server

import (
	"fmt"
	"hermes/cmd/config"
	"hermes/internals/api/routes"
	"hermes/internals/services"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	publisher := services.NewPublisher(&s.config.Rabbitmq)

	s.configure()

	// register routes
	routes.MessageRouter(api, publisher)
	routes.HealthCheck(api)

	port := s.config.Server.Port
	fmt.Printf("Server running on port %d\n", port)

	log.Fatal(s.app.Listen(fmt.Sprintf(":%d", port)))
}

func (s *Server) configure() {
	s.addCors()
	s.addRatelimiter()
	s.addLogger()
}

func (s *Server) addRatelimiter() {

	s.app.Use(limiter.New(limiter.Config{
		Max:               s.config.RateLimit.Max,
		Expiration:        time.Duration(s.config.RateLimit.Expiration) * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))
}

func (s *Server) addLogger() {
	s.app.Use(logger.New())
}

func (s *Server) addCors() {
	s.app.Use(cors.New())
}

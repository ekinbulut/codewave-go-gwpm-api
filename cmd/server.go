package main

import "fmt"

type Server struct {
	Config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Run() {
	fmt.Println("hello world")
}

package main

import (
	"github.com/Nimor111/grader/config"
	"github.com/Nimor111/grader/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.NewConfig(8000, "postgres", "password", "grader")
	s := server.New(config)

	if err := s.Run(); err != nil {
		log.Fatalf("Server failed to start!")
	}
}

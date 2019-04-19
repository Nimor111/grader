package server

import (
	"fmt"
	"github.com/Nimor111/grader/config"
	"github.com/Nimor111/grader/database"
	"github.com/Nimor111/grader/models/task"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Server struct {
	Config   config.Config
	Database database.Database
}

func New(cfg *config.Config) *Server {
	db := database.NewDB(*cfg)
	models := []interface{(*task.Task)(nil)}

	if err := db.CreateSchema([]interface{}{(*task.Task)(nil)}); err != nil {
		panic(fmt.Sprintf("Could not create database schema: %v", err))
	}

	return &Server{
		Config:   *cfg,
		Database: db,
	}
}

type defaultHandler struct{}

func (*defaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello docker!")
}

func (s *Server) Run() error {
	server := http.Server{
		Addr:    ":" + strconv.Itoa(s.Config.Port),
		Handler: &defaultHandler{},
	}

	log.Infof("Server started on port 8000")
	return server.ListenAndServe()
}

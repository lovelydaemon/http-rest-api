package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
  config *Config
  logger *logrus.Logger
  router *mux.Router
}

func New(config *Config) *APIServer {
  return &APIServer{
    config: config,
    logger: logrus.New(),
    router: mux.NewRouter(),
  }
}

func (s *APIServer) Start() error {
  err := s.ConfigureLogger()
  if err != nil {
    return err
  }

  s.ConfigureRouter()

  s.logger.Info("starting API server")

  return http.ListenAndServe(s.config.BindAddr, s.router) 
}

func (s *APIServer) ConfigureLogger() error {
  level, err := logrus.ParseLevel(s.config.LogLevel)
  if err != nil {
    return err
  }

  s.logger.SetLevel(level)

  return nil
}

func (s *APIServer) ConfigureRouter() {
  s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello")
  } 
}
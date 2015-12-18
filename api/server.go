package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/seguins/georges-manager/api/router/project"
	"net/http"
)

type Server struct {
	server *http.Server
	router *mux.Router
}

func NewServer() (*Server, error) {
	httpserver := &http.Server{
		Addr: ":8080",
	}
	s := &Server{
		server: httpserver,
		router: mux.NewRouter(),
	}

	return s, nil
}

func (s *Server) InitRouter() {
	project.InitRoutes(s.router)
}

func (s *Server) ServeAPI() error {
	s.server.Handler = s.router
	logrus.Infof("API listen on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

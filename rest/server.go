package rest

import (
	"net/http"
)

type server struct {
	globalMiddlewares Chain
	server            *http.ServeMux
}

func (s *server) Use(middlewares ...Middleware) {
	s.globalMiddlewares = s.globalMiddlewares.Append(middlewares...)
}

func (s *server) RegisterRoute(route Route) {
	s.server.HandleFunc(route.GetPattern(), route.Handler)
}

func NewServer() Server {
	mux := http.NewServeMux()

	return &server{
		server:            mux,
		globalMiddlewares: NewChain(),
	}
}

func (s *server) Start(addr string) error {
	return http.ListenAndServe(addr, s.globalMiddlewares.Then(s.server))
}

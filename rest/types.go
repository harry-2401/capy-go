package rest

import "net/http"

type (
	Route interface {
		Handler(w http.ResponseWriter, req *http.Request)
		GetPattern() string
	}

	// Middleware is a function that receives an http.Handler and returns an http.Handler.
	Middleware func(http.Handler) http.Handler

	Chain interface {
		Append(middlewares ...Middleware) Chain
		Prepend(middlewares ...Middleware) Chain
		Then(h http.Handler) http.Handler
		ThenFunc(fn http.HandlerFunc) http.Handler
		PrintChain() []Middleware
		AppendChain(chain Chain) Chain
		PrependChain(chain Chain) Chain
	}

	Server interface {
		Use(middlewares ...Middleware)
		RegisterRoute(route Route)
		Start(addr string) error
	}
)

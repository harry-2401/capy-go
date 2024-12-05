package rest

import "net/http"

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
	chain   Chain
}

func NewRoute(method, path string, handler http.HandlerFunc) Route {
	return &route{
		method:  method,
		path:    path,
		handler: handler,
		chain:   NewChain(),
	}
}

func (r *route) Handler(w http.ResponseWriter, req *http.Request) {
	r.chain.ThenFunc(r.handler).ServeHTTP(w, req)
}

func (r *route) SetChain(c Chain) {
	r.chain = c
}

func (r *route) GetPattern() string {
	return r.method + " " + r.path
}

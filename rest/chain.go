package rest

import (
	"net/http"
)

type chain struct {
	middlewares []Middleware
}

func NewChain(middlewares ...Middleware) Chain {
	return &chain{middlewares: middlewares}
}

func (c *chain) Append(middlewares ...Middleware) Chain {
	return &chain{middlewares: joinMiddlewares(c.middlewares, middlewares)}
}

func (c *chain) Prepend(middlewares ...Middleware) Chain {
	return &chain{middlewares: joinMiddlewares(middlewares, c.middlewares)}
}

func (c *chain) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range c.middlewares {
		h = c.middlewares[len(c.middlewares)-1-i](h)
	}

	return h
}

func (c *chain) ThenFunc(fn http.HandlerFunc) http.Handler {
	if fn == nil {
		return c.Then((nil))
	}

	return c.Then(fn)
}

func joinMiddlewares(a, b []Middleware) []Middleware {
	mids := make([]Middleware, 0, len(a)+len(b))
	mids = append(mids, a...)
	mids = append(mids, b...)
	return mids
}

func (c *chain) AppendChain(ch Chain) Chain {
	return &chain{middlewares: joinMiddlewares(c.middlewares, ch.PrintChain())}
}

func (c *chain) PrependChain(ch Chain) Chain {
	return &chain{middlewares: joinMiddlewares(ch.PrintChain(), c.middlewares)}
}

func (c *chain) PrintChain() []Middleware {
	return c.middlewares
}

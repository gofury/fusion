// package fusion provides chaing for both fasthttp RequestHandler and RequestHandler based Middleware
package fusion

import (
	"github.com/valyala/fasthttp"
	"github.com/gofury/furyrouter"
)

// type definition for Middleware that takes in a RequestHandler
type Middleware func(fasthttp.RequestHandler) fasthttp.RequestHandler

// Handlers acts as a simple function that
func Handlers(hs ...fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func (ctx *fasthttp.RequestCtx) {
		for _, h := range hs {
			if (ctx.Response.StatusCode() < 400) {
				h(ctx)
			}
		}
	}
}

// Abstraction for RequestHandler Middleware slice.
type Middlewares struct {
	middlewares []Middleware
}

// New creates a new Chain with given Middlewares
// Middlewares are only called upon a call to Then().
func New(ms ...Middleware) *Middlewares {
	return &Middlewares{ms}
}

// Handler chains the middlewares and returns the final http.Handler.
//     New(m1, m2, m3).Handle(h)
// is equivalent to:
//     m1(m2(m3(h)))
// When the request comes in, it will be passed m1 -> m2 -> m3 -> handler
// (assuming every middleware calls the following one).
//
// A chain can be safely reused by calling Handle() several times.
//     stdStack := fusion.New(ratelimitHandler, csrfHandler)
//     indexPipe = stdStack.Handler(indexHandler)
//     authPipe = stdStack.Handler(authHandler)
// For proper middleware, this should cause no problems.
//
// Handler() treats nil as furyrouter.New().Handler
func (m *Middlewares) Handler(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	if handler == nil {
		handler = furyrouter.New().Handler
	}
	return func (ctx *fasthttp.RequestCtx) {
		for i := range m.middlewares {
			handler = m.middlewares[len(m.middlewares)-1-i](handler)
		}
	}
}


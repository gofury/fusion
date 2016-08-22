package gofury

import (
	"github.com/valyala/fasthttp"
)

func Chain(hs ...fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func (ctx *fasthttp.RequestCtx) {
		for _, h := range hs {
			if (ctx.Response.StatusCode() < 400) {
				h(ctx)
			}
		}
	}
}
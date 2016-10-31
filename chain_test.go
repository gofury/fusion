package fusion

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)


func paramHandler(param string) fasthttp.RequestHandler {
	return func (ctx *fasthttp.RequestCtx) {
		if ctx.UserValue("params") == nil {
			ctx.SetUserValue("params", "")
		}
		ctx.SetUserValue("params", ctx.UserValue("params").(string) + param)
	}
}

func TestHandlersRunInOrder(t *testing.T) {
	// given
	h1 := paramHandler("1")
	h2 := paramHandler("2")
	h3 := paramHandler("3")
	ctx := &fasthttp.RequestCtx{}

	// when
	Handlers(h1, h2, h3)(ctx)

	// then
	assert.Equal(t, ctx.UserValue("params"), "123")
}

func tagMiddleware (tag string) Middleware {
	return func(h fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func (ctx *fasthttp.RequestCtx) {
			h(ctx)
			ctx.WriteString(tag)
		}
	}
}

var testHandler = func(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("\n"))
}

func TestMiddlewaresRunInOrder(t *testing.T) {
	// given
	m1 := tagMiddleware("1")
	m2 := tagMiddleware("2")
	m3 := tagMiddleware("3")
	ctx := &fasthttp.RequestCtx{}

	// when
	New(m1, m2, m3).Handler(testHandler)(ctx)

	// then
	assert.Equal(t, "123\n", string(ctx.Response.Body()))
}

//func TestStopOnResponseStatus(t *testing.T) {
//	// given
//	middleware1 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", "1")}
//	middleware2 := func(ctx *fasthttp.RequestCtx) {
//			ctx.SetUserValue("params", ctx.UserValue("params").(string) + "2")
//			ctx.SetStatusCode(fasthttp.StatusBadRequest)
//		}
//	middleware3 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", ctx.UserValue("params").(string) + "3")}
//	ctx := &fasthttp.RequestCtx{}
//
//	// when
//	Chain(middleware1, middleware2, middleware3)(ctx)
//
//	// then
//	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
//	assert.Equal(t, "12", ctx.UserValue("params"))
//}
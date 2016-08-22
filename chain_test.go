package gofury

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestHandlersRunInOrder(t *testing.T) {
	// given
	middleware1 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", "1")}
	middleware2 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", ctx.UserValue("params").(string) + "2")}
	middleware3 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", ctx.UserValue("params").(string) + "3")}
	ctx := &fasthttp.RequestCtx{}

	// when
	Chain(middleware1, middleware2, middleware3)(ctx)

	// then
	assert.Equal(t, ctx.UserValue("params"), "123")
}

func TestStopOnResponseStatus(t *testing.T) {
	// given
	middleware1 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", "1")}
	middleware2 := func(ctx *fasthttp.RequestCtx) {
			ctx.SetUserValue("params", ctx.UserValue("params").(string) + "2")
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
		}
	middleware3 := func(ctx *fasthttp.RequestCtx) {ctx.SetUserValue("params", ctx.UserValue("params").(string) + "3")}
	ctx := &fasthttp.RequestCtx{}

	// when
	Chain(middleware1, middleware2, middleware3)(ctx)

	// then
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
	assert.Equal(t, "12", ctx.UserValue("params"))
}
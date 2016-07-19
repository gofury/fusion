package fasthttpchain

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestNew(t *testing.T) {
	// given
	handlers := []fasthttp.RequestHandler{
		func(ctx *fasthttp.RequestCtx) {},
		func(ctx *fasthttp.RequestCtx) {},
		func(ctx *fasthttp.RequestCtx) {},
	}

	// when
	chain := RequestHandlerChain{handlers}

	// then
	assert.Equal(t, len(chain.requests), 3, "should have 2 request handlers")
}

func TestBuilder(t *testing.T) {
	// given
	builder := Builder()

	// when
	chain := builder.Append(func(ctx *fasthttp.RequestCtx) {}).
		Append(func(ctx *fasthttp.RequestCtx) {}, func(ctx *fasthttp.RequestCtx) {}).
		Build()

	// then
	assert.Equal(t, len(chain.requests), 3, "should have 3 request handlers")
}

func TestHandlersRunInOrder(t *testing.T) {
	// given
	handlers := []fasthttp.RequestHandler{
		func(ctx *fasthttp.RequestCtx) {
			ctx.SetUserValue("params", "1")
		},
		func(ctx *fasthttp.RequestCtx) {
			ctx.SetUserValue("params", ctx.UserValue("params").(string) + "2")
		},
		func(ctx *fasthttp.RequestCtx) {
			ctx.SetUserValue("params", ctx.UserValue("params").(string) + "3")
		},
	}
	chain := RequestHandlerChain{handlers}
	ctx := fasthttp.RequestCtx{}

	// when
	chain.HandlerChain(&ctx)

	// then
	assert.Equal(t, ctx.UserValue("params"), "123")
}
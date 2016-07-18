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
	chain := FastHttpChain{handlers}

	// then
	assert.Equal(t, len(chain.requests), 3, "should have 2 request handlers")
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
	chain := FastHttpChain{handlers}
	ctx := fasthttp.RequestCtx{}

	// when
	chain.ChainHandler(&ctx)

	// then
	assert.Equal(t, ctx.UserValue("params"), "123")
}
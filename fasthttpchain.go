package fasthttpchain

import (
	"github.com/valyala/fasthttp"
)

type RequestHandlerChain struct {
	handlers []fasthttp.RequestHandler
}

func New(hs ...fasthttp.RequestHandler) RequestHandlerChain {
	handlers := make([]fasthttp.RequestHandler, 0, len(hs))
	return RequestHandlerChain{append(handlers, hs...)}
}

func (chain *RequestHandlerChain) Handler(ctx *fasthttp.RequestCtx) {
	for i := range chain.handlers {
		chain.handlers[i](ctx)
	}
}

type FastHttpChainBuilder struct {
	chain RequestHandlerChain
}

func Builder() *FastHttpChainBuilder {
	return &FastHttpChainBuilder{RequestHandlerChain{}}
}

func (builder *FastHttpChainBuilder) Append(hs... fasthttp.RequestHandler) *FastHttpChainBuilder {
	builder.chain.handlers = append(builder.chain.handlers, hs...)
	return builder
}

func (builder *FastHttpChainBuilder) Build() RequestHandlerChain {
	return builder.chain
}
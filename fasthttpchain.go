package fasthttpchain

import (
	"github.com/valyala/fasthttp"
	//"github.com/buaazp/fasthttprouter"
)

type RequestHandlerChain struct {
	requests []fasthttp.RequestHandler
}

func (chain *RequestHandlerChain) HandlerChain(ctx *fasthttp.RequestCtx) {
	for i := range chain.requests {
		chain.requests[i](ctx)
	}
}

//func (chain *FastHttpChain) ChainRouter(ctx *fasthttp.RequestCtx, params *fasthttprouter.Params) {
//	for i := range chain.requests {
//		chain.requests[i](ctx, params)
//	}
//}

type FastHttpChainBuilder struct {
	chain RequestHandlerChain
}

func Builder() *FastHttpChainBuilder {
	return &FastHttpChainBuilder{RequestHandlerChain{}}
}

func (builder *FastHttpChainBuilder) Append(handlers... fasthttp.RequestHandler) *FastHttpChainBuilder {
	builder.chain.requests = append(builder.chain.requests, handlers...)
	return builder
}

func (builder *FastHttpChainBuilder) Build() RequestHandlerChain {
	return builder.chain
}
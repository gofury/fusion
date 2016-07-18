package fasthttpchain

import (
	"github.com/valyala/fasthttp"
)

type FastHttpChain struct {
	requests []fasthttp.RequestHandler
}

func (chain *FastHttpChain) ChainHandler(ctx *fasthttp.RequestCtx) {
	for i := range chain.requests {
		chain.requests[i](ctx)
		//h = c.constructors[len(c.constructors)-1-i](h)
	}
}

//func Builder() {
//	return &FastHttpChainBuilder{}
//}
//
//type FastHttpChainBuilder struct {
//	fastHttpChain FastHttpChain
//}
//
//func (builder *FastHttpChainBuilder) Append(hs ... fasthttp.RequestHandler) {
//	builder.fastHttpChain.requests = append(builder.fastHttpChain.requests, hs)
//}
//
//func (builder *FastHttpChainBuilder) Build(h fasthttp.RequestHandler) FastHttpChain {
//	return builder.fastHttpChain
//}
# FastHttpChain
FastHttp `RequestHandler` based Middleware chaining

[![Build Status](https://snap-ci.com/gofury/fasthttpchain/branch/master/build_image)](https://snap-ci.com/gofury/fasthttpchain/branch/master)
[![Coverage Status](https://coveralls.io/repos/github/gofury/fasthttpchain/badge.svg?branch=master)](https://coveralls.io/github/gofury/fasthttpchain?branch=master)

[GoDoc][godoc]

## Why
Chaining middleware calls allows you to turn:

    Middleware3(Middleware2(Middleware1(request)))
    
Into a more fluent format:

    New(Middleware1, Middleware2, Middleware3).Handler(request)

Because [fasthttp]() doesn't use `http.Handler` objects, but [RequestHandler][requestHandler] functions instead.
Therefore it is not possible to use any `http.Handler` based middleware chaining solutions such 
as [alice][alice] or [negroni][negroni]. 

However using `RequestHandlers` functions does [improves speed significantly][performance] wth zero memory allocations 
and also makes chaining code a lot simpler than object based chains.

## Usage
FastHttpChain is designed to be immutable. You can create a `RequestHandlerChain` either via constructor:

    chain := fasthttpchain.New(Middleware1, Middleware2, Middleware3)

or using a `Builder()`

    chain := fasthttpchain.Builder().Append(Middleware1).Append(Middleware2, Middleware3).Build()
    
Once you have `RequestHandlerChain`, you can pass in the `Handler` function to your `fasthttp` server. which 

    fasthttp.ListenAndServe("localhost:8080", chain.Handler)
    
The `Handler` function implements `RequestHandler` interface and will calls all handlers in the chain according to 
the order they are added.

[requestHandler]:   https://godoc.org/github.com/valyala/fasthttp#RequestHandler 
[performance]:      https://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp
[alice]:            https://github.com/justinas/alice/
[negroni]:          https://github.com/urfave/negroni
[godoc]:			https://godoc.org/github.com/gofury/fasthttpchain
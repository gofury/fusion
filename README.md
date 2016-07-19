# FastHttpChain
FastHttp Request Handler chaining

## Justifications

`fasthttp` doesn't support `http.Handler` objects but instead uses `fasthttp.RequestHandler` functions. 
Therefore it is not possible to use any `http.Handler` based middleware chaining solutions such as 
[alice](https://github.com/justinas/alice/) or [negroni](https://github.com/urfave/negroni). 

However, having `RequestHandlers` improves speed significantly wth zero `mallocs` and also makes chaining 
 middleware a lot easier.

## Usage

To ensure immutability, you can only create a `RequestHandlerChain` via constructor 

    chain := RequestHandlerChain{Middleware1, Middleware2, Middleware3}

or using the `Builder()`

    chain := fasthttpchain.Builder().Append(Middleware1).Append(Middleware2, Middleware3).Build()
    
Once you have `RequestHandlerChain`, you can pass in `HandlerChain` which implements `RequestHandler` to your 
`fasthttp` server.

    fasthttp.ListenAndServe("localhost:8080", chain.HandlerChain)
    

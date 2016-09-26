# Fusion

[![Build Status](https://snap-ci.com/gofury/fasthttpchain/branch/master/build_image)](https://snap-ci.com/gofury/fasthttpchain/branch/master)
[![Code Climate](https://codeclimate.com/github/gofury/fasthttpchain/badges/gpa.svg)](https://codeclimate.com/github/gofury/handlers)
[![GoDoc](http://godoc.org/github.com/gofury/handlers?status.png)](http://godoc.org/github.com/gofury/handlers)

Middleware chaining for [valyala/fasthttp][fasthttp]. Works for both:

 1. Traditional `Middlewares`: `func Middleware(h RequestHandler) RequestHandler`
 2. Pure [`RequestHandler`][requestHandler]: `func Handler(ctx *RequestCtx)`

This library is very much inspired by [alice]

## Usage

Sample code:

```
import (
    "github.com/valyala/fasthttp"
    "github.com/gofury/fusion"
)

func main() {
	router := furyrouter.New()
	
    router.GET("/admin", 
        // chain handlers is a simple function call
        fusion.Handlers(Handler1, Handler2, Handler2)
    )

    
    fasthttp.ListenAndServe(":8000", 
        // chain middleware requires a new struct
        fusion.New(Middleware1, Middleware2, Middleware3).Then(router)
    )
}
```

[fasthttp]:			https://github.com/valyala/fasthttp
[requestHandler]:   https://godoc.org/github.com/valyala/fasthttp#RequestHandler 
[godoc]:			https://godoc.org/github.com/gofury/fusion
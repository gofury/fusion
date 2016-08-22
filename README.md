# Gofury Handlers

[![Build Status](https://snap-ci.com/gofury/fasthttpchain/branch/master/build_image)](https://snap-ci.com/gofury/fasthttpchain/branch/master)
[![Code Climate](https://codeclimate.com/github/gofury/fasthttpchain/badges/gpa.svg)](https://codeclimate.com/github/gofury/handlers)
[![GoDoc](http://godoc.org/github.com/gofury/handlers?status.png)](http://godoc.org/github.com/gofury/handlers)

A collection of handlers (aka middleware) for use with [valyala/fasthttp][fasthttp] package and [RequestHandler][requestHandler]:

 * Chain - easy functional `RequestHandler` chaining.

## Usage

Chaining `RequestHandlers` is simple, everything is done via functions with no memory allocation:

```
import (
    "github.com/valyala/fasthttp"
    "github.com/gofury/handlers"
)

func main() {
	router := furyrouter.New()

	// chain 3 middlewares 
    router.GET("/admin", handlers.Chain(Middleware1, Middleware2, Middleware3))

    fasthttp.ListenAndServe(":8000", router)
}
```

## Why
using `RequestHandlers` functions does [improves speed significantly][performance] wth zero memory allocations 
and also makes chaining code a lot simpler than object based chains.

[fasthttp]:			https://github.com/valyala/fasthttp
[requestHandler]:   https://godoc.org/github.com/valyala/fasthttp#RequestHandler 
[performance]:      https://github.com/valyala/fasthttp#switching-from-nethttp-to-fasthttp
[godoc]:			https://godoc.org/github.com/gofury/fasthttpchain
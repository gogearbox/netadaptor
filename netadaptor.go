package netadaptor

import (
	"net/http"

	"github.com/gogearbox/gearbox"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// HTTPHandlerFunc wraps net/http handler func to gearbox handler
func HTTPHandlerFunc(h http.HandlerFunc) func(gearbox.Context) {
	return HTTPHandler(h)
}

// HTTPHandler wraps net/http handler to gearbox handler
func HTTPHandler(h http.Handler) func(gearbox.Context) {
	return func(ctx gearbox.Context) {
		handler := fasthttpadaptor.NewFastHTTPHandler(h)
		handler(ctx.Context())
	}
}

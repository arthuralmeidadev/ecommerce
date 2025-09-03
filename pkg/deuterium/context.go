package deuterium

import "net/http"

type primitiveContext interface {
	HttpResponseWriter() http.ResponseWriter
	HttpRequest() *http.Request
}

type wrappedContext interface {
	Request() *request
	Response() *response
	Next()
}

type Context interface {
	primitiveContext
	wrappedContext
}

type context struct {
	req   *request
	res   *response
	app   *app
	route *route
}

func (ctx *context) Request() *request {
	return ctx.req
}

func (ctx *context) Response() *response {
	return ctx.res
}

func (ctx *context) Next() {
	if ctx.route != nil {
		ctx.route.next(ctx)
	}
	ctx.app.next(ctx)
}

func (ctx *context) HttpResponseWriter() http.ResponseWriter {
	return ctx.res.w
}

func (ctx *context) HttpRequest() *http.Request {
	return ctx.req.r
}

package api

import (
	"net/http"
)

type primitiveContext interface {
	HttpResponseWriter() http.ResponseWriter
	HttpRequest() *http.Request
}

type wrappedContext interface {
	Request() *request
	Response() *response
}

type Context interface {
	primitiveContext
	wrappedContext
}

type context struct {
	req *request
	res *response
}

func (ctx *context) Request() *request {
	return ctx.req
}

func (ctx *context) Response() *response {
	return ctx.res
}

func (ctx *context) HttpResponseWriter() http.ResponseWriter {
	return ctx.res.w
}

func (ctx *context) HttpRequest() *http.Request {
	return ctx.req.r
}

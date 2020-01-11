package server

import (
	"context"
	"net/http"
)

type RequestContext interface {
	RoutePattern(req *http.Request) string
	PathParam(req *http.Request, key string) string
}

type contextKeyType struct{}

var contextKeyVal = contextKeyType{}

func WithRequestContext(ctx context.Context, reqCtx RequestContext) context.Context {
	return context.WithValue(ctx, contextKeyVal, reqCtx)
}

func fromContext(ctx context.Context) (RequestContext, bool) {
	reqCtx, ok := ctx.Value(contextKeyVal).(RequestContext)
	return reqCtx, ok
}

func RoutePattern(req *http.Request) string {
	reqCtx, ok := fromContext(req.Context())
	if !ok {
		return ""
	}
	return reqCtx.RoutePattern(req)
}

func PathParam(req *http.Request, key string) string {
	reqCtx, ok := fromContext(req.Context())
	if !ok {
		return ""
	}
	return reqCtx.PathParam(req, key)
}

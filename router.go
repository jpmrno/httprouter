package server

import "net/http"

type Router interface {
	RouterGroup
	http.Handler

	NotFound(handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	MethodNotAllowed(handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
}

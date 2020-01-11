package server

import "net/http"

type RouterGroup interface {
	Group(pattern string) RouterGroup

	Use(middlewares ...func(next http.Handler) http.Handler)

	Handle(method, pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Connect(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Delete(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Get(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Head(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Options(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Patch(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Post(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Put(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
	Trace(pattern string, handler http.Handler, middlewares ...func(next http.Handler) http.Handler)
}

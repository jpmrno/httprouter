package gin

import (
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
	"github.com/jpmrno/httprouter"
	"net/http"
)

type group struct {
	group *gin.RouterGroup
}

func wrapGroup(g *gin.RouterGroup) *group {
	return &group{group: g}
}

func (g group) Group(pattern string) server.RouterGroup {
	return wrapGroup(g.group.Group(pattern))
}

func (g group) Use(middlewares ...func(next http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		g.group.Use(adapter.Wrap(middleware))
	}
}

func (g group) Handle(method, pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.Handle(method, pattern, wrapHandlers(middlewares, handler))
}

func (g group) Connect(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.Handle(http.MethodConnect, pattern, wrapHandlers(middlewares, handler))
}

func (g group) Delete(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.DELETE(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Get(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.GET(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Head(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.HEAD(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Options(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.OPTIONS(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Patch(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.PATCH(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Post(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.POST(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Put(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.PUT(pattern, wrapHandlers(middlewares, handler))
}

func (g group) Trace(pattern string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	g.group.Handle(http.MethodTrace, pattern, wrapHandlers(middlewares, handler))
}

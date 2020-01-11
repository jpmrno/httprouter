package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jpmrno/httprouter"
	"net/http"
)

type router struct {
	*group
	engine *gin.Engine
}

func New() server.Router {
	return Wrap(gin.New())
}

func Wrap(engine *gin.Engine) server.Router {
	router := &router{
		group:  wrapGroup(&engine.RouterGroup),
		engine: engine,
	}
	router.engine.Use(func(ginCtx *gin.Context) {
		ctx := server.WithRequestContext(ginCtx.Request.Context(), wrapContext(ginCtx))
		ginCtx.Request = ginCtx.Request.WithContext(ctx)
	})
	return router
}

func (r router) NotFound(handler http.Handler, middlewares ...func(next http.Handler) http.Handler) {
	r.engine.NoRoute(wrapHandlers(middlewares, handler))
}

func (r router) MethodNotAllowed(handler http.Handler, middlewares ...func(next http.Handler) http.Handler) {
	r.engine.NoMethod(wrapHandlers(middlewares, handler))
}

func (r router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jpmrno/httprouter"
	"net/http"
)

type context struct {
	*gin.Context
}

func wrapContext(ginCtx *gin.Context) server.RequestContext {
	return context{Context: ginCtx}
}

func (c context) RoutePattern(req *http.Request) string {
	return c.FullPath()
}

func (c context) PathParam(req *http.Request, key string) string {
	return c.Param(key)
}

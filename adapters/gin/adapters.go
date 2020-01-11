package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/justinas/alice"
	"net/http"
)

func wrapHandlers(ms []func(http.Handler) http.Handler, h http.Handler) gin.HandlerFunc {
	var cs = make([]alice.Constructor, len(ms))
	for i, m := range ms {
		cs[i] = alice.Constructor(m)
	}
	return gin.WrapH(alice.New(cs...).Then(h))
}

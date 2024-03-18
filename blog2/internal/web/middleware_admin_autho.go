package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Impl) MiddlewareAdminAutho(ctx *gin.Context) {
	if !ctxGetAdmin(ctx) {
		t.P.RenderHTML(
			ctx,
			http.StatusNotFound,
			"page_error.html",
			PageError{
				ComponentCommonHead: ComponentCommonHead{},
				Message:             "Not found",
			},
		)
		ctx.Abort()
		return
	}
}

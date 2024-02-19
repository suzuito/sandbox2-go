package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (w *Impl) MiddlewareAdminAutho(ctx *gin.Context) {
	if !ctxGetAdmin(ctx) {
		w.P.RenderHTML(
			ctx,
			http.StatusNotFound,
			"page_error.html",
			viewmodel.PageError{
				ComponentCommonHead: viewmodel.ComponentCommonHead{},
				Message:             "Not found",
			},
		)
		ctx.Abort()
		return
	}
}

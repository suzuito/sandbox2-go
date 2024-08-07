package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Impl) PageNoRoute(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusNotFound,
		"page_error.html",
		PageError{
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			ComponentCommonHead: ComponentCommonHead{
				GoogleTagManagerID: t.GoogleTagManagerID,
			},
			Message: "Not found",
		},
	)
}

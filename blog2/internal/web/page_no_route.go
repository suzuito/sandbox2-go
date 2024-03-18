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
			ComponentCommonHead: ComponentCommonHead{},
			Message:             "Not found",
		},
	)
}

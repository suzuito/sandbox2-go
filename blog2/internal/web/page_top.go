package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageTop struct {
	ComponentCommonHead ComponentCommonHead
}

func (t *Impl) PageTop(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_top.html",
		PageTop{
			ComponentCommonHead: ComponentCommonHead{
				Title: "Login",
				Meta:  nil,
			},
		},
	)
}

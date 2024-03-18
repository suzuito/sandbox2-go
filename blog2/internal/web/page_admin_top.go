package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageAdminTop struct {
	ComponentCommonHead ComponentCommonHead
}

func (t *Impl) PageAdminTop(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_top.html",
		PageAdminTop{
			ComponentCommonHead: ComponentCommonHead{},
		},
	)
}

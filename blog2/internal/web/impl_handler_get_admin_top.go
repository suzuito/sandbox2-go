package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) GetAdminTop(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_top.html",
		viewmodel.PageAdminTop{
			ComponentCommonHead: viewmodel.ComponentCommonHead{},
		},
	)
}

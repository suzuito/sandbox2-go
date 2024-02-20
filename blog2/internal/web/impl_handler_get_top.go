package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) GetTop(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_top.html",
		viewmodel.PageTop{
			ComponentCommonHead: viewmodel.ComponentCommonHead{
				Title: "Login",
				Meta:  nil,
			},
		},
	)
}

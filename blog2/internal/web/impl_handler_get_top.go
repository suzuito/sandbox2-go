package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (w *Impl) GetTop(ctx *gin.Context) {
	w.P.RenderHTML(
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

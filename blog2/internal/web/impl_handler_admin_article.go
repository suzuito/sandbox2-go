package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) GetAdminArticle(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article.html",
		viewmodel.PageAdminArticle{
			ComponentCommonHead: viewmodel.ComponentCommonHead{},
			Article:             article,
		},
	)
}

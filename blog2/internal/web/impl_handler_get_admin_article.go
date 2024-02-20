package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) GetAdminArticle(ctx *gin.Context) {
	articleID := entity.ArticleID(ctx.Param("articleID"))
	dto, err := t.U.GetAdminArticle(ctx, articleID)
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			viewmodel.NewPageErrorUnknownError(),
		)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article.html",
		viewmodel.PageAdminArticle{
			ComponentCommonHead: viewmodel.ComponentCommonHead{},
			Article:             dto.Article,
		},
	)
}

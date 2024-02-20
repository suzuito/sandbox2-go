package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) GetAdminArticles(ctx *gin.Context) {
	query := entity.ArticleSearchQuery{}
	dto, err := t.U.GetAdminArticles(ctx, &query)
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			viewmodel.PageError{
				ComponentCommonHead: viewmodel.ComponentCommonHead{},
				Message:             "謎のエラーが発生した！",
			},
		)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_articles.html",
		viewmodel.PageAdminArticles{
			ComponentCommonHead: viewmodel.ComponentCommonHead{},
			Articles:            dto.Articles,
		},
	)
}

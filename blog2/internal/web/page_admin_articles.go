package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageAdminArticles struct {
	ComponentCommonHead ComponentCommonHead
	Articles            []*entity.Article
}

func (t *Impl) PageAdminArticles(ctx *gin.Context) {
	query := entity.ArticleSearchQuery{}
	dto, err := t.U.GetAdminArticles(ctx, &query)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			NewPageErrorUnknownError(),
		)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_articles.html",
		PageAdminArticles{
			ComponentCommonHead: ComponentCommonHead{},
			Articles:            dto.Articles,
		},
	)
}

func (t *Impl) PostAdminArticles(ctx *gin.Context) {
	t.L.ErrorContext(ctx, "TODO Check CSRF token")
	dto, err := t.U.PostAdminArticles(ctx)
	if err != nil {
		t.L.ErrorContext(ctx, "Failed to create article", "err", err)
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			NewPageErrorUnknownError(),
		)
		return
	}
	t.P.Redirect(
		ctx,
		http.StatusFound,
		fmt.Sprintf("/admin/articles/%s", dto.Article.ID),
	)
}
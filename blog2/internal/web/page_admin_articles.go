package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/cweb"
)

type PageAdminArticles struct {
	ComponentCommonHead        ComponentCommonHead
	ComponentHeader            ComponentHeader
	Articles                   []*entity.Article
	ComponentArticleListSearch ComponentArticleListSearch
	ComponentArticleListPager  ComponentArticleListPager
}

func (t *Impl) PageAdminArticles(ctx *gin.Context) {
	page := cweb.DefaultQueryAsInt(ctx, "page", 0)
	size := cweb.DefaultQueryAsInt(ctx, "size", 10)
	tagID := ctx.DefaultQuery("tag", "")
	var published *bool
	if _, exists := ctx.GetQuery("published"); exists {
		publishedValue := cweb.DefaultQueryAsBool(ctx, "published", false)
		published = &publishedValue
	}
	dto, err := t.U.GetAdminArticles(ctx, entity.TagID(tagID), page, size, published)
	if err != nil {
		t.L.Error("Failed to get admin articles", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_articles.html",
		PageAdminArticles{
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			ComponentCommonHead: ComponentCommonHead{
				GoogleTagManagerID: t.GoogleTagManagerID,
			},
			Articles:                   dto.Articles,
			ComponentArticleListSearch: ComponentArticleListSearch{},
			ComponentArticleListPager: ComponentArticleListPager{
				NextPage: dto.NextPage,
				PrevPage: dto.PrevPage,
			},
		},
	)
}

func (t *Impl) PostAdminArticles(ctx *gin.Context) {
	t.L.ErrorContext(ctx, "TODO Check CSRF token")
	dto, err := t.U.PostAdminArticles(ctx)
	if err != nil {
		t.L.ErrorContext(ctx, "Failed to create article", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.Redirect(
		ctx,
		http.StatusFound,
		fmt.Sprintf("/admin/articles/%s", dto.Article.ID),
	)
}

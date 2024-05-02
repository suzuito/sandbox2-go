package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/cweb"
)

type PageArticles struct {
	ComponentHeader           ComponentHeader
	ComponentCommonHead       ComponentCommonHead
	Articles                  []*entity.Article
	ComponentArticleListPager ComponentArticleListPager
	Breadcrumbs               Breadcrumbs
}

func (t *Impl) PageArticles(ctx *gin.Context) {
	tagID := ctx.DefaultQuery("tag", "")
	page := cweb.DefaultQueryAsInt(ctx, "page", 0)
	size := cweb.DefaultQueryAsInt(ctx, "size", 10)
	dto, err := t.U.PageArticles(ctx, entity.TagID(tagID), page, size)
	if err != nil {
		t.L.Error("Failed to get articles", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_articles.html",
		PageArticles{
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			ComponentCommonHead: ComponentCommonHead{
				GoogleTagManagerID: t.GoogleTagManagerID,
				Title:              fmt.Sprintf("%s - 記事一覧", SiteName),
				Meta:               nil,
			},
			Breadcrumbs: Breadcrumbs{
				{
					Path: "/",
					URL:  NewPageURL(t.SiteOrigin, "/"),
					Name: "トップページ",
				},
				{
					Name:   "記事一覧",
					NoLink: true,
				},
			},
			Articles: dto.Articles,
			ComponentArticleListPager: ComponentArticleListPager{
				NextPage: dto.NextPage,
				PrevPage: dto.PrevPage,
			},
		},
	)
}

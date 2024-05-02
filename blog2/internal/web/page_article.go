package web

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageArticle struct {
	ComponentHeader     ComponentHeader
	ComponentCommonHead ComponentCommonHead
	Breadcrumbs         Breadcrumbs
	Article             *entity.Article
	ArticleHTML         template.HTML
}

func (t *Impl) PageArticle(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	dto, err := t.U.PageArticle(ctx, article)
	if err != nil {
		t.L.Error("Failed to get articles", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_article.html",
		PageArticle{
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			ComponentCommonHead: ComponentCommonHead{
				GoogleTagManagerID: t.GoogleTagManagerID,
				Title:              fmt.Sprintf("%s - %s", SiteName, article.Title),
				Meta: &SiteMetaData{
					Canonical: NewPageURLFromContext(ctx, t.SiteOrigin),
					OGP: OGPData{
						Title:       article.Title,
						Description: "",
						Locale:      "ja_JP",
						Type:        "article",
						URL:         NewPageURLFromContext(ctx, t.SiteOrigin),
						SiteName:    SiteName,
						Image:       "",
					},
					LDJSON: []LDJSONData{
						{
							Context:          "https://schema.org",
							Type:             "Article",
							MainEntityOfPage: NewPageURLFromContext(ctx, t.SiteOrigin),
							Headline:         article.Title,
							Description:      "",
							DatePublished:    article.PublishedAt.Format(time.RFC3339),
							Author: &LDJSONDataAuthor{
								Type: "Person",
								Name: "suzuito",
							},
						},
					},
				},
			},
			Breadcrumbs: Breadcrumbs{
				{
					Path: "/",
					URL:  NewPageURL(t.SiteOrigin, "/"),
					Name: "トップページ",
				},
				{
					Path: "/articles",
					URL:  NewPageURL(t.SiteOrigin, "/articles"),
					Name: "記事一覧",
				},
				{
					Name:   article.Title,
					NoLink: true,
				},
			},
			Article:     article,
			ArticleHTML: template.HTML(dto.HTMLBody),
		},
	)
}

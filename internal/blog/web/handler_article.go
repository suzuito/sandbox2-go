package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

type responseGetArticles struct {
	Header      struct{}
	Articles    []entity.Article
	Meta        siteMetaData
	QueryTags   []string
	NextPage    int
	PrevPage    int
	HasNextPage bool
	HasPrevPage bool
}

func (t *ControllerImpl) GetArticles(ctx *gin.Context) {
	limit := 10
	query := struct {
		Page int    `form:"page"`
		Tags string `form:"tags"`
	}{}
	if err := ctx.BindQuery(&query); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	tags := strings.Split(query.Tags, ",")
	tagsFiltered := []string{}
	for _, tag := range tags {
		if tag == "" {
			continue
		}
		tagsFiltered = append(tagsFiltered, tag)
	}
	articles := []entity.Article{}
	offset := 0
	if query.Page > 0 {
		offset = query.Page * limit
	}
	hasNext := false
	if err := t.UC.SearchArticles(
		ctx,
		usecase.SearchArticlesQuery{
			Offset:    offset,
			Limit:     limit,
			Tags:      tagsFiltered,
			SortField: usecase.SearchArticlesQuerySortFieldDate,
			SortOrder: usecase.SortOrderDesc,
		},
		&articles,
		&hasNext,
	); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "pc_articles.html",
		Obj: responseGetArticles{
			Header:      struct{}{},
			Articles:    articles,
			NextPage:    query.Page + 1,
			PrevPage:    query.Page - 1,
			HasNextPage: hasNext,
			HasPrevPage: query.Page > 0,
			Meta: siteMetaData{
				OGP: ogpData{
					Title:       fmt.Sprintf("%s - 記事一覧", siteName),
					Description: "記事一覧",
					Locale:      "ja_JP",
					Type:        "website",
					URL:         getPageURL(ctx, t.Setting),
					SiteName:    siteName,
					Image:       "",
				},
				Canonical: getPageURL(ctx, t.Setting),
				LDJSON: []ldjsonData{
					{
						Context:          "https://schema.org",
						Type:             "WebSite",
						MainEntityOfPage: getPageURL(ctx, t.Setting),
						Headline:         siteName,
						Description:      "個人用ブログ",
					},
				},
			},
		},
	})
}

type responseGetArticle struct {
	Header      struct{}
	Article     *entity.Article
	ArticleHTML template.HTML
	Meta        siteMetaData
}

func (t *ControllerImpl) GetArticle(ctx *gin.Context) {
	html := bytes.NewBufferString("")
	article := ctxGetArticle(ctx)
	if err := t.RepositoryArticleHTML.GetArticle(ctx, article.ID, article.Version, html); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		return
	}
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusOK,
		Name: "pc_article.html",
		Obj: responseGetArticle{
			Header:      struct{}{},
			Article:     article,
			ArticleHTML: template.HTML(html.String()),
			Meta: siteMetaData{
				OGP: ogpData{
					Title:       fmt.Sprintf("%s - %s", siteName, article.Title),
					Description: article.Description,
					Locale:      "ja_JP",
					Type:        "article",
					URL:         getPageURL(ctx, t.Setting),
					SiteName:    siteName,
					Image:       "",
				},
				Canonical: getPageURL(ctx, t.Setting),
				LDJSON: []ldjsonData{
					{
						Context:          "https://schema.org",
						Type:             "Article",
						MainEntityOfPage: getPageURL(ctx, t.Setting),
						Headline:         article.Title,
						Description:      article.Description,
						DatePublished:    article.Date.Format(time.RFC3339),
						Author: ldjsonDataAuthor{
							Type: "Person",
							Name: "suzuito",
						},
					},
				},
			},
		},
	})
}

package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

type responseGetTop struct {
	Header   struct{}
	Articles []entity.Article
	Meta     siteMetaData
}

func (t *ControllerImpl) GetTop(ctx *gin.Context) {
	articles := []entity.Article{}
	hasNext := false
	if err := t.UC.SearchArticles(
		ctx,
		usecase.SearchArticlesQuery{
			Offset:    0,
			Limit:     10,
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
		Name: "pc_top.html",
		Obj: responseGetTop{
			Header:   struct{}{},
			Articles: articles,
			Meta: siteMetaData{
				OGP: ogpData{
					Title:       fmt.Sprintf("%s", siteName),
					Description: "個人用ブログ",
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

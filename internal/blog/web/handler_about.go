package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type responseCommon struct {
	Header struct{}
	Meta   siteMetaData
}

type responseGetAbout struct {
	responseCommon
	Articles []entity.Article
}

func (t *ControllerImpl) GetAbout(ctx *gin.Context) {
	t.Presenters.Response(
		ctx,
		PresenterArgHTML{
			Code: http.StatusOK,
			Name: "pc_about.html",
			Obj: responseGetAbout{
				responseCommon: responseCommon{
					Header: struct{}{},
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
			},
		},
	)
}

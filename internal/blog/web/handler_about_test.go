package web

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAbout(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com:1234"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/about", ctrl.GetAbout)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/about", nil)
				return r
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "pc_about.html",
				Obj: responseGetAbout{
					responseCommon: responseCommon{
						Header: struct{}{},
						Meta: siteMetaData{
							OGP: ogpData{
								Title:       "otiuzu pages",
								Description: "個人用ブログ",
								Locale:      "ja_JP",
								Type:        "website",
								URL:         "https://www.example.com:1234/about",
								SiteName:    "otiuzu pages",
								Image:       "",
							},
							Canonical: "https://www.example.com:1234/about",
							LDJSON: []ldjsonData{
								{
									Context:          "https://schema.org",
									Type:             "WebSite",
									MainEntityOfPage: "https://www.example.com:1234/about",
									Headline:         "otiuzu pages",
									Description:      "個人用ブログ",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(
			tC.String(),
			testWebEndpoint(setting, &tC, setRouterFunc),
		)
	}
}

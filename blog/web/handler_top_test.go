package web

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/blog/usecase"
)

func TestGetTop(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/", ctrl.GetTop)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				articles := []entity.Article{}
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     10,
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).
					SetArg(2, articles).
					SetArg(3, false).
					Return(nil).
					Times(1)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "pc_top.html",
				Obj: responseGetTop{
					Header:   struct{}{},
					Articles: []entity.Article{},
					Meta: siteMetaData{
						OGP: ogpData{
							Title:       "otiuzu pages",
							Description: "個人用ブログ",
							Locale:      "ja_JP",
							Type:        "website",
							URL:         "https://www.example.com/",
							SiteName:    "otiuzu pages",
							Image:       "",
						},
						Canonical: "https://www.example.com/",
						LDJSON: []ldjsonData{
							{
								Context:          "https://schema.org",
								Type:             "WebSite",
								MainEntityOfPage: "https://www.example.com/",
								Headline:         "otiuzu pages",
								Description:      "個人用ブログ",
							},
						},
					},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				err := errors.New("dummy error")
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     10,
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).
					Return(err).
					Times(1)
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: errors.New("dummy error"),
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

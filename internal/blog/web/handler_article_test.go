package web

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

func TestGetArticles(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/articles", ctrl.GetArticles)
	}
	testCases := []testCaseWebEndpoint{
		{
			desc: "no query",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     10,
						Tags:      []string{},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).
					SetArg(2, []entity.Article{}).
					SetArg(3, false).
					Return(nil).
					Times(1)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "pc_articles.html",
				Obj: responseGetArticles{
					Header:      struct{}{},
					Articles:    []entity.Article{},
					NextPage:    1,
					PrevPage:    -1,
					HasNextPage: false,
					HasPrevPage: false,
					Meta: siteMetaData{
						OGP: ogpData{
							Title:       "otiuzu pages - 記事一覧",
							Description: "記事一覧",
							Locale:      "ja_JP",
							Type:        "website",
							URL:         "https://www.example.com/articles",
							SiteName:    "otiuzu pages",
							Image:       "",
						},
						Canonical: "https://www.example.com/articles",
						LDJSON: []ldjsonData{
							{
								Context:          "https://schema.org",
								Type:             "WebSite",
								MainEntityOfPage: "https://www.example.com/articles",
								Headline:         "otiuzu pages",
								Description:      "個人用ブログ",
							},
						},
					},
				},
			},
		},
		{
			desc: "query",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles?page=10&tags=tag1,tag2", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset: 100,
						Limit:  10,
						Tags: []string{
							"tag1",
							"tag2",
						},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).
					SetArg(2, []entity.Article{}).
					SetArg(3, true).
					Return(nil).
					Times(1)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "pc_articles.html",
				Obj: responseGetArticles{
					Header:      struct{}{},
					Articles:    []entity.Article{},
					NextPage:    11,
					PrevPage:    9,
					HasNextPage: true,
					HasPrevPage: true,
					Meta: siteMetaData{
						OGP: ogpData{
							Title:       "otiuzu pages - 記事一覧",
							Description: "記事一覧",
							Locale:      "ja_JP",
							Type:        "website",
							URL:         "https://www.example.com/articles",
							SiteName:    "otiuzu pages",
							Image:       "",
						},
						Canonical: "https://www.example.com/articles",
						LDJSON: []ldjsonData{
							{
								Context:          "https://schema.org",
								Type:             "WebSite",
								MainEntityOfPage: "https://www.example.com/articles",
								Headline:         "otiuzu pages",
								Description:      "個人用ブログ",
							},
						},
					},
				},
			},
		},
		{
			desc: "bind query error",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles?page=hoge", nil)
				return r
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: errors.New("strconv.ParseInt: parsing \"hoge\": invalid syntax"),
			},
		},
		{
			desc: "search articles error",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     10,
						Tags:      []string{},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).
					Return(errors.New("dummy error")).
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

func TestGetArticle(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/articles/:articleID", ctrl.GetArticle)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a01", nil)
				return r
			},
			beforeMiddleware: func(ctx *gin.Context) {
				ctxSetArticle(ctx, &entity.Article{
					ID:          "a01",
					Version:     1,
					Title:       "ダミータイトル",
					Description: "ダミータイトルです！",
				})
			},
			setUp: func(
				m *mocks,
			) {
				b := bytes.NewBufferString("dummy html")
				m.RepositoryArticleHTML.EXPECT().GetArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					int32(1),
					gomock.Any(),
				).SetArg(3, *b).Times(1)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "pc_article.html",
				Obj: responseGetArticle{
					Header: struct{}{},
					Article: &entity.Article{
						ID:          "a01",
						Version:     1,
						Title:       "ダミータイトル",
						Description: "ダミータイトルです！",
					},
					ArticleHTML: template.HTML("dummy html"),
					Meta: siteMetaData{
						OGP: ogpData{
							Title:       "otiuzu pages - ダミータイトル",
							Description: "ダミータイトルです！",
							Locale:      "ja_JP",
							Type:        "article",
							URL:         "https://www.example.com/articles/a01",
							SiteName:    "otiuzu pages",
							Image:       "",
						},
						Canonical: "https://www.example.com/articles/a01",
						LDJSON: []ldjsonData{
							{
								Context:          "https://schema.org",
								Type:             "Article",
								MainEntityOfPage: "https://www.example.com/articles/a01",
								Headline:         "ダミータイトル",
								Description:      "ダミータイトルです！",
								DatePublished:    "0001-01-01T00:00:00Z",
								Author: ldjsonDataAuthor{
									Type: "Person",
									Name: "suzuito",
								},
							},
						},
					},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/articles/a01", nil)
				return r
			},
			beforeMiddleware: func(ctx *gin.Context) {
				ctxSetArticle(ctx, &entity.Article{
					ID:          "a01",
					Version:     1,
					Title:       "ダミータイトル",
					Description: "ダミータイトルです！",
				})
			},
			setUp: func(
				m *mocks,
			) {
				m.RepositoryArticleHTML.EXPECT().GetArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					int32(1),
					gomock.Any(),
				).Return(errors.New("dummy error")).Times(1)
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

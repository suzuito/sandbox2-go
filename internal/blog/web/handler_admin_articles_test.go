package web

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func TestGetAdminArticlesByID(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("", ctrl.GetAdminArticlesByID)
	}
	testCases := []testCaseWebEndpoint{
		{
			desc: "ok",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			beforeMiddleware: func(ctx *gin.Context) {
				ctxSetArticle(ctx, &entity.Article{
					ID:      "a01",
					Version: 1,
				})
			},
			setUp: func(m *mocks) {
				b := bytes.NewBufferString("dummy html")
				m.RepositoryArticleHTML.EXPECT().GetArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					int32(1),
					gomock.Any(),
				).SetArg(3, *b).Times(1)
				m.RepositoryArticle.EXPECT().GetArticlesByID(
					gomock.Any(),
					entity.ArticleID("a01"),
					gomock.Any(),
				).SetArg(2, []entity.Article{
					{ID: "a01", Version: 1, ArticleSource: entity.ArticleSource{ID: "as01", Version: "v1"}},
					{ID: "a01", Version: 2, ArticleSource: entity.ArticleSource{ID: "as02", Version: "v1"}},
					{ID: "a01", Version: 3, ArticleSource: entity.ArticleSource{ID: "as03", Version: "v1"}},
				}).Times(1)
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"main",
					entity.ArticleSourceID("as01"),
				).Return([]entity.ArticleSource{
					{
						ID:      entity.ArticleSourceID("as01"),
						Version: "v1",
						Meta: entity.ArticleSourceMeta{
							URL: "http://www.example.com/as01.html?v=v1",
						},
					},
				}, nil).Times(1)
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"main",
					entity.ArticleSourceID("as02"),
				).Return([]entity.ArticleSource{
					{
						ID:      entity.ArticleSourceID("as02"),
						Version: "v1",
						Meta: entity.ArticleSourceMeta{
							URL: "http://www.example.com/as02.html?v=v1",
						},
					},
				}, nil).Times(1)
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"main",
					entity.ArticleSourceID("as03"),
				).Return([]entity.ArticleSource{
					{
						ID:      entity.ArticleSourceID("as03"),
						Version: "v1",
						Meta: entity.ArticleSourceMeta{
							URL: "http://www.example.com/as03.html?v=v1",
						},
					},
				}, nil).Times(1)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_articles_by_id.html",
				Obj: responseGetAdminArticlesByID{
					responseCommon: responseCommon{
						Header: struct{}{},
						Meta:   siteMetaData{},
					},
					Article: entity.Article{
						ID:      "a01",
						Version: 1,
					},
					ArticleHTML: template.HTML("dummy html"),
					ArticlesByID: []entity.Article{
						{
							ID:      "a01",
							Version: 1,
							ArticleSource: entity.ArticleSource{
								ID:      "as01",
								Version: "v1",
							},
						},
						{
							ID:      "a01",
							Version: 2,
							ArticleSource: entity.ArticleSource{
								ID:      "as02",
								Version: "v1",
							},
						},
						{
							ID:      "a01",
							Version: 3,
							ArticleSource: entity.ArticleSource{
								ID:      "as03",
								Version: "v1",
							},
						},
					},
					MapArticleSourceToArticleSourceVersions: map[entity.ArticleSourceID][]entity.ArticleSource{
						"as01": []entity.ArticleSource{
							{
								ID:      "as01",
								Version: "v1",
								Meta: entity.ArticleSourceMeta{
									URL: "http://www.example.com/as01.html?v=v1",
								},
							},
						},
						"as02": []entity.ArticleSource{
							{
								ID:      "as02",
								Version: "v1",
								Meta: entity.ArticleSourceMeta{
									URL: "http://www.example.com/as02.html?v=v1",
								},
							},
						},
						"as03": []entity.ArticleSource{
							{
								ID:      "as03",
								Version: "v1",
								Meta: entity.ArticleSourceMeta{
									URL: "http://www.example.com/as03.html?v=v1",
								},
							},
						},
					},
					MapArticleSourceVersionToArticleVersion: map[string]int32{
						"v1": 3,
					},
				},
			},
		},
		{
			desc: "err1",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			beforeMiddleware: func(ctx *gin.Context) {
				ctxSetArticle(ctx, &entity.Article{
					ID:      "a01",
					Version: 1,
				})
			},
			setUp: func(m *mocks) {
				m.RepositoryArticleHTML.EXPECT().GetArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					int32(1),
					gomock.Any(),
				).Return(fmt.Errorf("dummy error")).Times(1)
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: fmt.Errorf("dummy error"),
			},
		},
		{
			desc: "err2",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			beforeMiddleware: func(ctx *gin.Context) {
				ctxSetArticle(ctx, &entity.Article{
					ID:      "a01",
					Version: 1,
				})
			},
			setUp: func(m *mocks) {
				b := bytes.NewBufferString("dummy html")
				m.RepositoryArticleHTML.EXPECT().GetArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					int32(1),
					gomock.Any(),
				).SetArg(3, *b).Times(1)
				m.RepositoryArticle.EXPECT().GetArticlesByID(
					gomock.Any(),
					entity.ArticleID("a01"),
					gomock.Any(),
				).Return(fmt.Errorf("dummy error")).Times(1)
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: fmt.Errorf("dummy error"),
			},
		},
		{
			desc: "err3",
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			beforeMiddleware: func(ctx *gin.Context) {
				ctxSetArticle(ctx, &entity.Article{
					ID:      "a01",
					Version: 1,
				})
			},
			setUp: func(m *mocks) {
				b := bytes.NewBufferString("dummy html")
				m.RepositoryArticleHTML.EXPECT().GetArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					int32(1),
					gomock.Any(),
				).SetArg(3, *b).Times(1)
				m.RepositoryArticle.EXPECT().GetArticlesByID(
					gomock.Any(),
					entity.ArticleID("a01"),
					gomock.Any(),
				).SetArg(2, []entity.Article{
					{ID: "a01", Version: 1, ArticleSource: entity.ArticleSource{ID: "as01", Version: "v1"}},
					{ID: "a01", Version: 2, ArticleSource: entity.ArticleSource{ID: "as02", Version: "v1"}},
					{ID: "a01", Version: 3, ArticleSource: entity.ArticleSource{ID: "as03", Version: "v1"}},
				}).Times(1)
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"main",
					entity.ArticleSourceID("as01"),
				).Return(nil, fmt.Errorf("dummy error")).Times(1)
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: fmt.Errorf("dummy error"),
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

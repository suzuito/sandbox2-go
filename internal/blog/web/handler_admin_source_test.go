package web

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func TestGetAdminSearchSource(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("", ctrl.GetAdminSearchSource)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/?q=hoge", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.RepositoryArticleSource.EXPECT().SearchArticleSources(
					gomock.Any(),
					"hoge",
					gomock.Any(),
				)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_sources_search.html",
				Obj: gin.H{
					"ArticleSources": []*entity.ArticleSource{},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/?q=hoge", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.RepositoryArticleSource.EXPECT().SearchArticleSources(
					gomock.Any(),
					"hoge",
					gomock.Any(),
				).Return(errors.New("dummy error"))
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

func TestGetAdminSourcesByID(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/:articleSourceID", ctrl.GetAdminSourcesByID)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/as01?branch=hoge", nil)
				return r
			},
			setUp: func(m *mocks) {
				articleSources := []entity.ArticleSource{
					{ID: "as01", Version: "v01"},
					{ID: "as01", Version: "v02"},
					{ID: "as01", Version: "v03"},
				}
				articles := []entity.Article{
					{ID: "a01", Version: 901, ArticleSource: entity.ArticleSource{ID: "as01", Version: "v01"}},
				}
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"hoge",
					entity.ArticleSourceID("as01"),
				).Return(
					articleSources,
					nil,
				)
				m.RepositoryArticle.EXPECT().GetArticlesByArticleSourceID(
					gomock.Any(),
					entity.ArticleSourceID("as01"),
					gomock.Any(),
				).SetArg(2, articles)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_sources_by_id.html",
				Obj: gin.H{
					"ArticleSources": []entity.ArticleSource{
						{ID: "as01", Version: "v01"},
						{ID: "as01", Version: "v02"},
						{ID: "as01", Version: "v03"},
					},
					"MapArticleSourceVersionToArticle": map[string]*entity.Article{
						"v01": &entity.Article{
							ID: "a01", Version: 901, ArticleSource: entity.ArticleSource{ID: "as01", Version: "v01"},
						},
					},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/as01?branch=hoge", nil)
				return r
			},
			setUp: func(m *mocks) {
				articleSources := []entity.ArticleSource{
					{ID: "as01", Version: "v01"},
					{ID: "as01", Version: "v02"},
					{ID: "as01", Version: "v03"},
				}
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"hoge",
					entity.ArticleSourceID("as01"),
				).Return(
					articleSources,
					errors.New("dummy error"),
				)
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: errors.New("dummy error"),
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/as01?branch=hoge", nil)
				return r
			},
			setUp: func(m *mocks) {
				articleSources := []entity.ArticleSource{
					{ID: "as01", Version: "v01"},
					{ID: "as01", Version: "v02"},
					{ID: "as01", Version: "v03"},
				}
				articles := []entity.Article{}
				m.RepositoryArticleSource.EXPECT().GetVersions(
					gomock.Any(),
					"hoge",
					entity.ArticleSourceID("as01"),
				).Return(
					articleSources,
					nil,
				)
				m.RepositoryArticle.EXPECT().GetArticlesByArticleSourceID(
					gomock.Any(),
					entity.ArticleSourceID("as01"),
					gomock.Any(),
				).SetArg(2, articles).Return(errors.New("dummy error"))
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

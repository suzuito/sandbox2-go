package web

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/blog/entity"
)

func TestMiddlewareGetArticle(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.Use(ctrl.MiddlewareGetArticle)
		e.GET("/:articleID", func(ctx *gin.Context) {
			ctrl.Presenters.Response(ctx, PresenterArgJSON{
				Code: http.StatusOK,
				Obj:  *ctxGetArticle(ctx),
			})
		})
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/a01", nil)
				return r
			},
			setUp: func(m *mocks) {
				article := entity.Article{ID: "a01"}
				m.RepositoryArticle.EXPECT().GetLatestArticle(
					gomock.Any(),
					entity.ArticleID("a01"),
					gomock.Any(),
				).SetArg(2, article).Times(1)
			},
			expectedPresenterArg: PresenterArgJSON{
				Code: http.StatusOK,
				Obj:  entity.Article{ID: "a01"},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/a01", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.RepositoryArticle.EXPECT().GetLatestArticle(
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
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/a01?version=1", nil)
				return r
			},
			setUp: func(m *mocks) {
				article := entity.Article{ID: "a01"}
				m.RepositoryArticle.EXPECT().GetArticleByPrimaryKey(
					gomock.Any(),
					entity.ArticlePrimaryKey{ArticleID: "a01", Version: 1},
					gomock.Any(),
				).SetArg(2, article).Times(1)
			},
			expectedPresenterArg: PresenterArgJSON{
				Code: http.StatusOK,
				Obj:  entity.Article{ID: "a01"},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/a01?version=1", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.RepositoryArticle.EXPECT().GetArticleByPrimaryKey(
					gomock.Any(),
					entity.ArticlePrimaryKey{ArticleID: "a01", Version: 1},
					gomock.Any(),
				).Return(fmt.Errorf("dummy error")).Times(1)
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: fmt.Errorf("dummy error"),
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/a01?version=a", nil)
				return r
			},
			setUp: func(m *mocks) {
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: fmt.Errorf(`strconv.ParseInt: parsing "a": invalid syntax`),
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

package web

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/blog/usecase"
)

func TestGetAdmin(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("", ctrl.GetAdmin)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     1000,
						Tags:      []string{},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				)
				m.RepositoryArticleSource.EXPECT().GetBranches(gomock.Any()).Return(
					[]string{"br1", "br2"},
					nil,
				)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_top.html",
				Obj: gin.H{
					"Branches": []string{"br1", "br2"},
					"Articles": []entity.Article{},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/?tags=t1,t2&page=10", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    10000,
						Limit:     1000,
						Tags:      []string{"t1", "t2"},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				)
				m.RepositoryArticleSource.EXPECT().GetBranches(gomock.Any()).Return(
					[]string{"br1", "br2"},
					nil,
				)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_top.html",
				Obj: gin.H{
					"Branches": []string{"br1", "br2"},
					"Articles": []entity.Article{},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/?page=a", nil)
				return r
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: errors.New(`strconv.ParseInt: parsing "a": invalid syntax`),
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     1000,
						Tags:      []string{},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				).Return(fmt.Errorf("dummy error"))
			},
			expectedPresenterArg: PresenterArgStandardError{
				Err: fmt.Errorf("dummy error"),
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().SearchArticles(
					gomock.Any(),
					usecase.SearchArticlesQuery{
						Offset:    0,
						Limit:     1000,
						Tags:      []string{},
						SortField: usecase.SearchArticlesQuerySortFieldDate,
						SortOrder: usecase.SortOrderDesc,
					},
					gomock.Any(),
					gomock.Any(),
				)
				m.RepositoryArticleSource.EXPECT().GetBranches(gomock.Any()).Return(
					[]string{"br1", "br2"},
					fmt.Errorf("dummy error"),
				)
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

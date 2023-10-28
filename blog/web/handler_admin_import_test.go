package web

import (
	"errors"
	"html/template"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/blog/entity"
)

func TestGetAdminImport(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("import/:articleSourceID/:articleSourceVersion", ctrl.GetAdminImport)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/import/as01/v01", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().GenerateArticleHTML(gomock.Any(), entity.ArticleSourceID("as01"), "v01").
					Return(
						&entity.Article{
							ID:      "as01returned",
							Version: 999,
						},
						"dummy html body",
						nil,
					)
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_import.html",
				Obj: gin.H{
					"ArticleSourceID": entity.ArticleSourceID("as01"),
					"Version":         "v01",
					"Article": &entity.Article{
						ID:      "as01returned",
						Version: 999,
					},
					"ArticleHTML": template.HTML("dummy html body"),
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/import/as01/v01", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().GenerateArticleHTML(gomock.Any(), entity.ArticleSourceID("as01"), "v01").
					Return(
						nil,
						"",
						errors.New("dummy error"),
					)
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

func TestGetAdminImportSuccess(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("import/:articleSourceID/:articleSourceVersion", ctrl.GetAdminImportSuccess)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/import/as01/v01?articleID=a01&articleVersion=999", nil)
				return r
			},
			setUp: func(m *mocks) {
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_result_success.html",
				Obj: gin.H{
					"ArticleID":            entity.ArticleID("a01"),
					"ArticleVersion":       "999",
					"ArticleSourceID":      entity.ArticleSourceID("as01"),
					"ArticleSourceVersion": "v01",
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

func TestGetAdminImportError(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("import/:articleSourceID/:articleSourceVersion", ctrl.GetAdminImportError)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/import/as01/v01?articleID=a01&articleVersion=999", nil)
				return r
			},
			setUp: func(m *mocks) {
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusOK,
				Name: "admin_result_error.html",
				Obj: gin.H{
					"ArticleSourceID":      entity.ArticleSourceID("as01"),
					"ArticleSourceVersion": "v01",
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

func TestPostAdminImport(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.POST("import/:articleSourceID/:articleSourceVersion", ctrl.PostAdminImport)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodPost, "/import/as01/v01", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().UploadArticle(gomock.Any(), entity.ArticleSourceID("as01"), "v01").
					Return(&entity.Article{
						ID:      "as01returned",
						Version: 999,
					}, nil)
			},
			expectedPresenterArg: PresenterArgRedirect{
				Code:     http.StatusFound,
				Location: "/admin/import/as01/v01/success?articleID=as01returned&articleVersion=999",
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodPost, "/import/as01/v01", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().UploadArticle(gomock.Any(), entity.ArticleSourceID("as01"), "v01").
					Return(nil, errors.New("dummy error"))
			},
			expectedPresenterArg: PresenterArgRedirect{
				Code:     http.StatusFound,
				Location: "/admin/import/as01/v01/error",
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

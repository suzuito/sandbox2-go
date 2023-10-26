package web

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestPostAdminExecuteImportSources(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.POST("", ctrl.PostAdminExecuteImportSources)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodPost, "/", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().UploadAllArticles(gomock.Any(), "main")
			},
			expectedPresenterArg: PresenterArgRedirect{
				Code:     http.StatusFound,
				Location: "/admin",
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodPost, "/", nil)
				return r
			},
			setUp: func(m *mocks) {
				m.UC.EXPECT().UploadAllArticles(gomock.Any(), "main").Return(fmt.Errorf("dummy error"))
			},
			expectedPresenterArg: PresenterArgRedirect{
				Code:     http.StatusFound,
				Location: "/admin",
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

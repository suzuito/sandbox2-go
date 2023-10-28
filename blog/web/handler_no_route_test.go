package web

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNoRoute(t *testing.T) {
	setting := newWebSetting()
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("", ctrl.NoRoute)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/", nil)
				return r
			},
			expectedPresenterArg: PresenterArgHTML{
				Code: http.StatusNotFound,
				Name: "pc_error.html",
				Obj: responseError{
					Title:   "404 - ページが存在しません",
					Message: "Not found",
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

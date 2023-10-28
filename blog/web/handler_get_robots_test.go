package web

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetRobotsTXT(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/robots.txt", ctrl.GetRobots)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/robots.txt", nil)
				return r
			},
			expectedPresenterArg: PresenterArgString{
				Code: http.StatusOK,
				Body: "Sitemap: https://www.example.com/sitemap.xml",
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

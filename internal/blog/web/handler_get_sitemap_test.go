package web

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

func TestGetSitemap(t *testing.T) {
	setting := newWebSetting()
	setting.SiteOrigin = "https://www.example.com"
	setRouterFunc := func(e *gin.Engine, ctrl *ControllerImpl) {
		e.GET("/sitemap.xml", ctrl.GetSitemap)
	}
	testCases := []testCaseWebEndpoint{
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/sitemap.xml", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				urls := usecase.XMLURLSet{
					URLs: []usecase.XMLURL{
						{Loc: "https://www.example.com/a1.html", Lastmod: "2023-01-02"},
						{Loc: "https://www.example.com/a2.html", Lastmod: "2023-01-03"},
					},
				}
				m.UC.EXPECT().GenerateSitemap(gomock.Any(), "https://www.example.com", gomock.Any()).
					SetArg(2, urls).
					Return(nil).
					Times(1)

			},
			expectedPresenterArg: PresenterArgXML{
				Code: http.StatusOK,
				Obj: usecase.XMLURLSet{
					URLs: []usecase.XMLURL{
						{
							Loc:     "https://www.example.com/a1.html",
							Lastmod: "2023-01-02",
						},
						{
							Loc:     "https://www.example.com/a2.html",
							Lastmod: "2023-01-03",
						},
					},
				},
			},
		},
		{
			inputRequestFunc: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "/sitemap.xml", nil)
				return r
			},
			setUp: func(
				m *mocks,
			) {
				m.UC.EXPECT().GenerateSitemap(gomock.Any(), "https://www.example.com", gomock.Any()).
					Return(fmt.Errorf("dummy error")).
					Times(1)
			},
			expectedPresenterArg: PresenterArgXML{
				Code: http.StatusOK,
				Obj:  usecase.XMLURLSet{},
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

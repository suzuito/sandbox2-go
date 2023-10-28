package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *ControllerImpl) GetRobots(ctx *gin.Context) {
	t.Presenters.Response(ctx, PresenterArgString{
		Code: http.StatusOK,
		Body: fmt.Sprintf(
			`Sitemap: %s/sitemap.xml`,
			t.Setting.SiteOrigin,
		),
	})
}

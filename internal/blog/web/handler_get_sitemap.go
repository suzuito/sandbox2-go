package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

func (t *ControllerImpl) GetSitemap(ctx *gin.Context) {
	urls := usecase.XMLURLSet{}
	t.UC.GenerateSitemap(ctx, t.Setting.SiteOrigin, &urls)
	t.Presenters.Response(ctx, PresenterArgXML{
		Code: http.StatusOK,
		Obj:  urls,
	})
}

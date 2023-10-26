package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
)

func (t *ControllerImpl) PostAdminExecuteImportSources(ctx *gin.Context) {
	if err := t.UC.UploadAllArticles(ctx, "main"); err != nil {
		clog.L.Errorf(ctx, "%+v", err)
	}
	t.Presenters.Response(ctx, PresenterArgRedirect{
		Code:     http.StatusFound,
		Location: "/admin",
	})
}

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *ControllerImpl) NoRoute(ctx *gin.Context) {
	t.Presenters.Response(ctx, PresenterArgHTML{
		Code: http.StatusNotFound,
		Name: "pc_error.html",
		Obj: responseError{
			Title:   "404 - ページが存在しません",
			Message: "Not found",
		},
	})
}

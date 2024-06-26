package web

import "github.com/gin-gonic/gin"

type PhotoStudioID string

func (t *Impl) MiddlewarePhotoStudio(ctx *gin.Context) {
	photoStudioID := ctx.Param("photoStudioID")
	t.L.Debug("", "photoStudioID", photoStudioID)
	ctx.Next()
}

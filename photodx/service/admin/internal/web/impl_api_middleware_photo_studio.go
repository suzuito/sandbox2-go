package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) APIMiddlewarePhotoStudio(
	ctx *gin.Context,
) {
	photoStudioID := ctx.Param("photoStudioID")
	dto, err := t.U.APIMiddlewarePhotoStudio(
		ctx,
		entity.PhotoStudioID(photoStudioID),
	)
	if err != nil {
		var noEntryError *repository.NoEntryError
		if errors.As(err, &noEntryError) {
			t.P.JSON(ctx, http.StatusNotFound, common_web.ResponseError{
				Message: fmt.Sprintf("PhotoStudioID '%s' is not found", photoStudioID),
			})
		} else {
			t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
				Message: "internal server error",
			})
		}
		ctx.Abort()
		return
	}
	common_web.CtxSet(ctx, common_web.CtxPhotoStudio, dto.PhotoStudio)
	ctx.Next()
}

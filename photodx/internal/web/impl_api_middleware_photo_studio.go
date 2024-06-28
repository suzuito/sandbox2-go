package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service/repository"
)

func (t *Impl) APIMiddlewarePhotoStudio(ctx *gin.Context) {
	photoStudioID := ctx.Param("photoStudioID")
	dto, err := t.U.APIMiddlewarePhotoStudio(ctx, entity.PhotoStudioID(photoStudioID))
	if err != nil {
		var noEntryError *repository.NoEntryError
		if errors.As(err, &noEntryError) {
			t.P.JSON(ctx, http.StatusNotFound, ResponseError{
				Message: fmt.Sprintf("PhotoStudioID '%s' is not found", photoStudioID),
			})
		} else {
			t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{
				Message: "internal server error",
			})
		}
		ctx.Abort()
		return
	}
	ctxSet(ctx, ctxPhotoStudio, dto.PhotoStudio)
	ctx.Next()
}

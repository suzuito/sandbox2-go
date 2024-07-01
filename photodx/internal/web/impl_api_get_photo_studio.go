package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) APIGetPhotoStudio(ctx *gin.Context) {
	t.P.JSON(ctx, http.StatusOK, ctxGet[*entity.PhotoStudio](ctx, ctxPhotoStudio))
}

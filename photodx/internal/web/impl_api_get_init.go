package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) APIGetInit(ctx *gin.Context) {
	dto, err := t.U.APIGetInit(ctx, ctxGet[entity.Principal](ctx, ctxPrincipal))
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{
			Message: "internal server error",
		})
		return
	}
	t.P.JSON(
		ctx,
		http.StatusOK,
		struct {
			PhotoStudio       *entity.PhotoStudio       `json:"photoStudio"`
			PhotoStudioMember *entity.PhotoStudioMember `json:"photoStudioMember"`
		}{
			PhotoStudio:       dto.PhotoStudio,
			PhotoStudioMember: dto.PhotoStudioMember,
		},
	)
}

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) APIGetInit(ctx *gin.Context) {
	principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
	dto, err := t.U.APIGetInit(ctx, principal)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
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

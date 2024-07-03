package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) AuthPostRefresh(ctx *gin.Context) {
	principal := common_web.CtxGet[entity.AdminPrincipalRefreshToken](ctx, common_web.CtxPrincipalRefreshToken)
	dto, err := t.U.AuthPostRefresh(ctx, principal.GetPhotoStudioMemberID())
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
			Message: "internal server error",
		})
		return
	}
	t.P.JSON(ctx, http.StatusCreated, struct {
		AccessToken string `json:"accessToken"`
	}{
		AccessToken: dto.AccessToken,
	})
}

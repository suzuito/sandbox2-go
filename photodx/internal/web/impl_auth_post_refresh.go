package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) AuthPostRefresh(ctx *gin.Context) {
	principal := ctxGet[entity.PrincipalRefreshToken](ctx, ctxPrincipalRefreshToken)
	dto, err := t.U.AuthPostRefresh(ctx, principal.GetPhotoStudioMemberID())
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{
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

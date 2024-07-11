package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/auth/internal/businesslogic"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) AuthPostLogin(ctx *gin.Context) {
	message := struct {
		PhotoStudioID entity.PhotoStudioID `json:"photoStudioId"`
		Email         string               `json:"email"`
		Password      string               `json:"password"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		t.P.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
			Message: "bad request",
		})
		return
	}
	dto, err := t.U.AuthPostLogin(ctx, message.PhotoStudioID, message.Email, message.Password)
	if err != nil {
		if errors.Is(err, businesslogic.ErrPasswordMismatch) {
			t.P.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
				Message: "email or password is invalid",
			})
			return
		}
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
			Message: "email or password is invalid",
		})
		return
	}
	t.P.JSON(ctx, http.StatusOK, struct {
		RefreshToken string `json:"refreshToken"`
	}{
		RefreshToken: dto.RefreshToken,
	})
}

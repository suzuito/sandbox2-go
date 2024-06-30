package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
	"github.com/suzuito/sandbox2-go/photodx/internal/usecase/service"
)

func (t *Impl) AuthPostLogin(ctx *gin.Context) {
	message := struct {
		PhotoStudioID entity.PhotoStudioID `json:"photoStudioId"`
		Email         string               `json:"email"`
		Password      string               `json:"password"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		t.P.JSON(ctx, http.StatusBadRequest, ResponseError{
			Message: "bad request",
		})
		return
	}
	dto, err := t.U.AuthPostLogin(ctx, message.PhotoStudioID, message.Email, message.Password)
	if err != nil {
		if errors.Is(err, service.ErrPasswordMismatch) {
			t.P.JSON(ctx, http.StatusBadRequest, ResponseError{
				Message: "email or password is invalid",
			})
			return
		}
		t.P.JSON(ctx, http.StatusInternalServerError, ResponseError{
			Message: "internal server error",
		})
		return
	}
	t.P.JSON(ctx, http.StatusOK, struct {
		RefreshToken string `json:"refreshToken"`
	}{
		RefreshToken: dto.RefreshToken,
	})
}

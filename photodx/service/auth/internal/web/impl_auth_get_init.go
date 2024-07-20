package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/repository"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) AuthGetInit(ctx *gin.Context) {
	principal := common_web.CtxGetAdminPrincipalAccessToken(ctx)
	dto, err := t.U.AuthGetInit(ctx, principal)
	if err != nil {
		var noEntryError *repository.NoEntryError
		if errors.As(err, &noEntryError) {
			t.P.JSON(ctx, http.StatusNotFound, common_web.ResponseError{
				Message: "not found",
			})
			return
		}
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
			Message: "internal server error",
		})
		return
	}
	t.P.JSON(ctx, http.StatusOK, dto)
}

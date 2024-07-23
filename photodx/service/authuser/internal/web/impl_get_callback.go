package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity/oauth2loginflow"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) GetCallback(ctx *gin.Context) {
	error := ctx.Query("error")
	errorDescription := ctx.Query("error_description")
	if error != "" {
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
			Message: fmt.Sprintf("internal server error because %s", errorDescription),
		})
		return
	}
	code := ctx.Query("code")
	state := ctx.Query("state")
	dto, err := t.U.GetCallback(
		ctx,
		code,
		oauth2loginflow.StateCode(state),
	)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
			Message: fmt.Sprintf("internal server error because %s", err.Error()),
		})
		return
	}
	redirectURL := t.FrontBaseURL
	query := redirectURL.Query()
	query.Set("refreshToken", dto.RefreshToken)
	redirectURL.RawQuery = query.Encode()
	fmt.Println(t.FrontBaseURL.String())
	fmt.Println(redirectURL.String())
	ctx.Redirect(http.StatusFound, redirectURL.String())
}

package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) GetAuthorizeLine(
	ctx *gin.Context,
) {
	dto, err := t.U.GetAuthorizeURLLINE(
		ctx,
		&t.OAuth2RedirectURL,
	)
	if err != nil {
		t.P.JSON(ctx, http.StatusInternalServerError, common_web.ResponseError{
			Message: "internal server error",
		})
		return
	}
	ctx.Redirect(http.StatusFound, dto.AuthorizeURL.String())
}

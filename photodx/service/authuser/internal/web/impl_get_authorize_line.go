package web

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) GetAuthorizeLine(
	ctx *gin.Context,
) {
	callbackURLString := ctx.DefaultQuery("callback", "")
	if callbackURLString == "" {
		t.P.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
			Message: "callback is empty",
		})
		return
	}
	callbackURL, err := url.Parse(callbackURLString)
	if err != nil {
		t.P.JSON(ctx, http.StatusBadRequest, common_web.ResponseError{
			Message: "callback is invalid",
		})
		return
	}
	dto, err := t.U.GetAuthorizeURLLINE(
		ctx,
		callbackURL,
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

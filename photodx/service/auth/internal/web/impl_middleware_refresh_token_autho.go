package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) MiddlewareRefreshTokenAutho(ctx *gin.Context) {
	principal := common_web.CtxGet[entity.AdminPrincipalRefreshToken](ctx, common_web.CtxPrincipalRefreshToken)
	if principal == nil {
		t.P.JSON(
			ctx,
			http.StatusUnauthorized,
			common_web.ResponseError{
				Message: "unauthorized",
			},
		)
		ctx.Abort()
		return
	}
	ctx.Next()
}

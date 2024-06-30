package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) MiddlewareRefreshTokenAutho(ctx *gin.Context) {
	principal := ctxGet[entity.PrincipalRefreshToken](ctx, ctxPrincipalRefreshToken)
	if principal == nil {
		t.P.JSON(
			ctx,
			http.StatusUnauthorized,
			ResponseError{
				Message: "unauthorized",
			},
		)
		ctx.Abort()
		return
	}
	ctx.Next()
}

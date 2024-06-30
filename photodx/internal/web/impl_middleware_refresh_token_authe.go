package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/internal/entity"
)

func (t *Impl) MiddlewareRefreshTokenAuthe(ctx *gin.Context) {
	authorizationHeaderString := ctx.GetHeader("Authorization")
	authorizationHeaderParts := strings.Fields(authorizationHeaderString)
	if len(authorizationHeaderParts) != 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
		ctx.Next()
		return
	}
	refreshToken := authorizationHeaderParts[1]
	dto, err := t.U.MiddlewareRefreshTokenAuthe(ctx, refreshToken)
	if err != nil {
		t.L.Warn("refreshToken authe is failed", "err", err)
		ctx.Next()
		return
	}
	ctxSet[entity.PrincipalRefreshToken](ctx, ctxPrincipalRefreshToken, dto.Principal)
	ctx.Next()
}

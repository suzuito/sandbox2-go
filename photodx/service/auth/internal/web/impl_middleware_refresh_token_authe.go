package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
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
	common_web.CtxSet(ctx, common_web.CtxPrincipalRefreshToken, dto.Principal)
	ctx.Next()
}

package web

import (
	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) MiddlewareRefreshTokenAuthe(ctx *gin.Context) {
	refreshToken := common_web.ExtractBearerToken(ctx)
	if refreshToken == "" {
		ctx.Next()
		return
	}
	dto, err := t.U.MiddlewareRefreshTokenAuthe(ctx, refreshToken)
	if err != nil {
		t.L.Warn("refreshToken authe is failed", "err", err)
		ctx.Next()
		return
	}
	common_web.CtxSet(ctx, common_web.CtxPrincipalRefreshToken, dto.Principal)
	ctx.Next()
}

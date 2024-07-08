package web

import (
	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) MiddlewareAccessTokenAuthe(ctx *gin.Context) {
	accessToken := common_web.ExtractBearerToken(ctx)
	if accessToken == "" {
		ctx.Next()
		return
	}
	dto, err := t.U.MiddlewareAccessTokenAuthe(ctx, accessToken)
	if err != nil {
		t.L.Warn("accessToken authe is failed", "err", err)
		ctx.Next()
		return
	}
	common_web.CtxSet(ctx, common_web.CtxPrincipal, dto.Principal)
	ctx.Next()
}

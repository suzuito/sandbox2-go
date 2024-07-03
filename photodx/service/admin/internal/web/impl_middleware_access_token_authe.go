package web

import (
	"strings"

	"github.com/gin-gonic/gin"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

func (t *Impl) MiddlewareAccessTokenAuthe(ctx *gin.Context) {
	authorizationHeaderString := ctx.GetHeader("Authorization")
	authorizationHeaderParts := strings.Fields(authorizationHeaderString)
	if len(authorizationHeaderParts) != 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
		ctx.Next()
		return
	}
	accessToken := authorizationHeaderParts[1]
	dto, err := t.U.MiddlewareAccessTokenAuthe(ctx, accessToken)
	if err != nil {
		t.L.Warn("accessToken authe is failed", "err", err)
		ctx.Next()
		return
	}
	common_web.CtxSet(ctx, common_web.CtxPrincipal, dto.Principal)
	ctx.Next()
}

package web

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (t *Impl) APIMiddlewareAuthAuthe(ctx *gin.Context) {
	authorizationHeaderString := ctx.GetHeader("Authorization")
	authorizationHeaderParts := strings.Fields(authorizationHeaderString)
	if len(authorizationHeaderParts) != 2 || strings.ToLower(authorizationHeaderParts[0]) != "bearer" {
		ctx.Next()
		return
	}
	accessToken := authorizationHeaderParts[1]
	dto, err := t.U.APIMiddlewareAuthAuthe(ctx, accessToken)
	if err != nil {
		t.L.Warn("authe is failed", "err", err)
		ctx.Next()
		return
	}
	ctxSet(ctx, ctxPrincipal, dto.Principal)
	ctx.Next()
}

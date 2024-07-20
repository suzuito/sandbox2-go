package web

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/authuser/internal/entity"
	common_web "github.com/suzuito/sandbox2-go/photodx/service/common/pkg/web"
)

// UserPrincipalRefreshToken
func CtxSetUserPrincipalRefreshToken(ctx *gin.Context, principal entity.UserPrincipalRefreshToken) {
	common_web.CtxSet(ctx, "UserPrincipalRefreshToken", principal)
}
func CtxGetUserPrincipalRefreshToken(ctx *gin.Context) entity.UserPrincipalRefreshToken {
	return common_web.CtxGet[entity.UserPrincipalRefreshToken](ctx, "UserPrincipalRefreshToken")
}

package web

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/photodx/service/common/pkg/entity"
)

type CtxKey string

const (
	CtxPhotoStudio           CtxKey = "PhotoStudio"
	CtxPrincipalRefreshToken CtxKey = "PrincipalRefreshToken"
)

func CtxSet[T any](ctx *gin.Context, k CtxKey, v T) {
	ctx.Set(string(k), v)
}

func CtxGet[T any](ctx *gin.Context, k CtxKey) T {
	var zero T // Tがポインタである場合、zeroはnilとなる
	v, ok := ctx.Get(string(k))
	if !ok {
		return zero
	}
	vv, ok := v.(T)
	if !ok {
		return zero
	}
	return vv
}

// AdminPrincipalAccessToken
func CtxSetAdminPrincipalAccessToken(ctx *gin.Context, principal entity.AdminPrincipalAccessToken) {
	CtxSet(ctx, "AdminPrincipalAccessToken", principal)
}
func CtxGetAdminPrincipalAccessToken(ctx *gin.Context) entity.AdminPrincipalAccessToken {
	return CtxGet[entity.AdminPrincipalAccessToken](ctx, "AdminPrincipalAccessToken")
}

// UserPrincipalAccessToken
func CtxSetUserPrincipalAccessToken(ctx *gin.Context, principal entity.UserPrincipalAccessToken) {
	CtxSet(ctx, "UserPrincipalAccessToken", principal)
}
func CtxGetUserPrincipalAccessToken(ctx *gin.Context) entity.UserPrincipalAccessToken {
	return CtxGet[entity.UserPrincipalAccessToken](ctx, "UserPrincipalAccessToken")
}

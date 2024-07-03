package web

import (
	"github.com/gin-gonic/gin"
)

type CtxKey string

const (
	CtxPrincipal             CtxKey = "Principal"
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

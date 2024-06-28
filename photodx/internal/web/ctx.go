package web

import (
	"github.com/gin-gonic/gin"
)

type ctxKey string

const (
	ctxPrincipal   ctxKey = "Principal"
	ctxPhotoStudio ctxKey = "PhotoStudio"
)

func ctxSet[T any](ctx *gin.Context, k ctxKey, v T) {
	ctx.Set(string(k), v)
}

func ctxGet[T any](ctx *gin.Context, k ctxKey) T {
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

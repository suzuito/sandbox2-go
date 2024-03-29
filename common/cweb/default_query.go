package cweb

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// DefaultQueryAsInt ...
func DefaultQueryAsInt(ctx *gin.Context, name string, dflt int) int {
	v := ctx.Query(name)
	if v == "" {
		return dflt
	}
	vv, err := strconv.Atoi(v)
	if err != nil {
		return dflt
	}
	return vv
}

// DefaultQueryAsInt64 ...
func DefaultQueryAsInt64(ctx *gin.Context, name string, dflt int64) int64 {
	v := ctx.Query(name)
	if v == "" {
		return dflt
	}
	vv, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return dflt
	}
	return vv
}

// DefaultQueryAsFloat32 ...
func DefaultQueryAsFloat32(ctx *gin.Context, name string, dflt float32) float32 {
	v := ctx.Query(name)
	if v == "" {
		return dflt
	}
	vv, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return dflt
	}
	return float32(vv)
}

// DefaultQueryAsFloat64 ...
func DefaultQueryAsFloat64(ctx *gin.Context, name string, dflt float64) float64 {
	v := ctx.Query(name)
	if v == "" {
		return dflt
	}
	vv, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return dflt
	}
	return vv
}

// DefaultQueryAsBool ...
func DefaultQueryAsBool(ctx *gin.Context, name string, dflt bool) bool {
	v := ctx.Query(name)
	if v == "" {
		return dflt
	}
	vv, err := strconv.ParseBool(v)
	if err != nil {
		return dflt
	}
	return vv
}

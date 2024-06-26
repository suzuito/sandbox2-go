package presenter

import "github.com/gin-gonic/gin"

type Presenter interface {
	JSON(
		ctx *gin.Context,
		code int,
		obj any,
	)
}

type Impl struct{}

func (t *Impl) JSON(
	ctx *gin.Context,
	code int,
	obj any,
) {
	ctx.JSON(code, obj)
}

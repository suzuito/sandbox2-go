package web

import "github.com/gin-gonic/gin"

type PresenterJSON struct {
}

func (t *PresenterJSON) Response(
	ctx *gin.Context,
	arg *PresenterArgJSON,
) {
	ctx.JSON(arg.Code, arg.Obj)
}

type PresenterArgJSON struct {
	Code int
	Obj  any
}

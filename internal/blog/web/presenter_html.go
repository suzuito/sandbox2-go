package web

import (
	"github.com/gin-gonic/gin"
)

type PresenterHTML struct {
}

func (t *PresenterHTML) Response(
	ctx *gin.Context,
	arg *PresenterArgHTML,
) {
	ctx.Set("code", arg.Code)
	ctx.HTML(arg.Code, arg.Name, arg.Obj)
}

type PresenterArgHTML struct {
	Code int
	Name string
	Obj  any
}

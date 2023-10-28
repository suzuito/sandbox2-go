package web

import "github.com/gin-gonic/gin"

type PresenterXML struct {
}

func (t *PresenterXML) Response(
	ctx *gin.Context,
	arg *PresenterArgXML,
) {
	ctx.XML(arg.Code, arg.Obj)
}

type PresenterArgXML struct {
	Code int
	Obj  any
}

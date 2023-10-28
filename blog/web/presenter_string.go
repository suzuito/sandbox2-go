package web

import "github.com/gin-gonic/gin"

type PresenterString struct {
}

func (t *PresenterString) Response(
	ctx *gin.Context,
	arg *PresenterArgString,
) {
	ctx.String(arg.Code, arg.Body)
}

type PresenterArgString struct {
	Code int
	Body string
}

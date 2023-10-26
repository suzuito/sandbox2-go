package web

import "github.com/gin-gonic/gin"

type PresenterRedirect struct {
}

func (t *PresenterRedirect) Response(
	ctx *gin.Context,
	arg *PresenterArgRedirect,
) {
	ctx.Redirect(arg.Code, arg.Location)
}

type PresenterArgRedirect struct {
	Code     int
	Location string
}

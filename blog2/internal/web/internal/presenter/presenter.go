package presenter

import "github.com/gin-gonic/gin"

type Presenter interface {
	RenderJSON(
		ctx *gin.Context,
		code int,
		obj any,
	)
	RenderHTML(
		ctx *gin.Context,
		code int,
		name string,
		obj any,
	)
	Redirect(
		ctx *gin.Context,
		code int,
		url string,
	)
}

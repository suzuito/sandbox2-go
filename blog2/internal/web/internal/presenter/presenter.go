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
	RenderString(
		ctx *gin.Context,
		code int,
		message string,
	)
	Redirect(
		ctx *gin.Context,
		code int,
		url string,
	)
}

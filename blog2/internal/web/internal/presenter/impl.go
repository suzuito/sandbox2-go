package presenter

import "github.com/gin-gonic/gin"

type Impl struct{}

func (t *Impl) RenderJSON(
	ctx *gin.Context,
	code int,
	obj any,
) {
	ctx.JSON(code, obj)
}

func (t *Impl) RenderHTML(
	ctx *gin.Context,
	code int,
	name string,
	obj any,
) {
	ctx.HTML(code, name, obj)
}

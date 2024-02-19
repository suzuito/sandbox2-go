package web

import (
	"github.com/gin-gonic/gin"
)

func (w *Impl) MiddlewareAdminAuthe(ctx *gin.Context) {
	token, err := ctx.Cookie("admin_auth_token")
	if err != nil {
		return
	}
	if token != w.AdminToken {
		return
	}
	ctxSetAdmin(ctx)
}

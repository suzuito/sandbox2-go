package web

import (
	"github.com/gin-gonic/gin"
)

func (t *Impl) MiddlewareAdminAuthe(ctx *gin.Context) {
	token, err := ctx.Cookie("admin_auth_token")
	if err != nil {
		return
	}
	if token != t.AdminToken {
		return
	}
	ctxSetAdmin(ctx)
}

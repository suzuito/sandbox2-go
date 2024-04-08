package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageAdminFiles struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
}

func (t *Impl) GetAdminFiles(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_files.html",
		PageAdminFiles{},
	)
}

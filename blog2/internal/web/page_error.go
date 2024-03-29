package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageError struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
	Message             string
}

func (t *Impl) Render4XXError(ctx *gin.Context, code int, message string) {
	t.P.RenderHTML(
		ctx,
		code,
		"page_error.html",
		PageError{
			ComponentCommonHead: ComponentCommonHead{},
			ComponentHeader: ComponentHeader{
				ctxGetAdmin(ctx),
			},
			Message: message,
		},
	)
}

func (t *Impl) RenderUnknownError(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusInternalServerError,
		"page_error.html",
		PageError{
			ComponentCommonHead: ComponentCommonHead{},
			ComponentHeader: ComponentHeader{
				ctxGetAdmin(ctx),
			},
			Message: "謎のエラーが発生した！",
		},
	)
}

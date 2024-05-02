package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageTop struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
}

func (t *Impl) PageTop(ctx *gin.Context) {
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_top.html",
		PageTop{
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			ComponentCommonHead: ComponentCommonHead{
				Title:              fmt.Sprintf("%s - トップ", SiteName),
				Meta:               nil,
				GoogleTagManagerID: t.GoogleTagManagerID,
			},
		},
	)
}

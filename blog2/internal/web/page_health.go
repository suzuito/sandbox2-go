package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Impl) PageHealth(ctx *gin.Context) {
	t.P.RenderJSON(ctx, http.StatusOK, gin.H{"status": "ok"})
}

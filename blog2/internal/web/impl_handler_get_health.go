package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (w *Impl) GetHealth(ctx *gin.Context) {
	w.P.RenderJSON(ctx, http.StatusOK, gin.H{"status": "ok"})
}

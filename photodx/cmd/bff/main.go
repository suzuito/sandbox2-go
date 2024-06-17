package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hoge", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hoge")
	})
	engine.Run()
}
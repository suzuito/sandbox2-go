package cweb

import "github.com/gin-gonic/gin"

// Google search engineにインデックスされないようにする
// https://developers.google.com/search/docs/crawling-indexing/block-indexing
func MiddlewareXRobotsTag() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("X-Robots-Tag", "noindex")
		ctx.Next()
	}
}

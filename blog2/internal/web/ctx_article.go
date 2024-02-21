package web

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

var ctxKeyArticle = "article"

func ctxSetArticle(ctx *gin.Context, article *entity.Article) {
	ctx.Set(ctxKeyArticle, article)
}

func ctxGetArticle(ctx *gin.Context) *entity.Article {
	v, ok := ctx.Get(ctxKeyArticle)
	if !ok {
		return nil
	}
	vv, ok := v.(*entity.Article)
	if !ok {
		return nil
	}
	return vv
}

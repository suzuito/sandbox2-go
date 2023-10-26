package web

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func (t *ControllerImpl) MiddlewareGetLatestArticle(ctx *gin.Context) {
	article := entity.Article{}
	if err := t.RepositoryArticle.GetLatestArticle(
		ctx,
		entity.ArticleID(ctx.Param("articleID")),
		&article,
	); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		ctx.Abort()
		return
	}
	ctxSetArticle(ctx, &article)
}

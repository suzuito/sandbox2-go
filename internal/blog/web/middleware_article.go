package web

import (
	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func (t *ControllerImpl) MiddlewareGetArticle(ctx *gin.Context) {
	query := struct {
		Version int32 `form:"version"`
	}{}
	if err := ctx.BindQuery(&query); err != nil {
		t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
		ctx.Abort()
		return
	}
	article := entity.Article{}
	if query.Version > 0 {
		if err := t.RepositoryArticle.GetArticleByPrimaryKey(
			ctx,
			entity.ArticlePrimaryKey{
				ArticleID: entity.ArticleID(ctx.Param("articleID")),
				Version:   query.Version,
			},
			&article,
		); err != nil {
			t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
			ctx.Abort()
			return
		}
	} else {
		if err := t.RepositoryArticle.GetLatestArticle(
			ctx,
			entity.ArticleID(ctx.Param("articleID")),
			&article,
		); err != nil {
			t.Presenters.Response(ctx, PresenterArgStandardError{Err: err})
			ctx.Abort()
			return
		}
	}
	ctxSetArticle(ctx, &article)
}

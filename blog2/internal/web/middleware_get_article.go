package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase/pkg/serviceerror"
)

func (t *Impl) GetMiddlewareGetArticle(publishedOnly bool) func(*gin.Context) {
	return func(ctx *gin.Context) {
		articleID := entity.ArticleID(ctx.Param("articleID"))
		dto, err := t.U.MiddlewareGetArticle(ctx, articleID, publishedOnly)
		if err != nil {
			if errors.As(err, &serviceerror.PtrNotFoundEntityError) {
				t.Render4XXError(ctx, http.StatusNotFound, "not found")
				ctx.Abort()
				return
			}
			t.L.Error("", "err", err)
			t.RenderUnknownError(ctx)
			ctx.Abort()
			return
		}
		ctxSetArticle(ctx, dto.Article)
	}
}

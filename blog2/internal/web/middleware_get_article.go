package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase"
)

func (t *Impl) MiddlewareGetArticle(ctx *gin.Context) {
	articleID := entity.ArticleID(ctx.Param("articleID"))
	dto, err := t.U.MiddlewareGetArticle(ctx, articleID)
	if err != nil {
		if errors.As(err, &usecase.PtrNotFoundEntityError) {
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

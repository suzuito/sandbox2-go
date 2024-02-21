package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) MiddlewareGetArticle(ctx *gin.Context) {
	articleID := entity.ArticleID(ctx.Param("articleID"))
	dto, err := t.U.MiddlewareGetArticle(ctx, articleID)
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			viewmodel.NewPageErrorUnknownError(),
		)
		ctx.Abort()
		return
	}
	ctxSetArticle(ctx, dto.Article)
}

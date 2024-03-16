package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/web/viewmodel"
)

func (t *Impl) GetAdminArticleTags(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	dto, err := t.U.GetAdminArticleTags(ctx, article)
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			viewmodel.NewPageErrorUnknownError(),
		)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article_tags.html",
		viewmodel.PageAdminArticleTags{
			ComponentCommonHead: viewmodel.ComponentCommonHead{},
			JsEnv:               viewmodel.PageAdminArticleTagsJsEnv{},
			Article:             article,
			Tags:                dto.Tags,
		},
	)
}

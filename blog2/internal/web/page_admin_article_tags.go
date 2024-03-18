package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageAdminArticleTags struct {
	ComponentCommonHead ComponentCommonHead
	Article             *entity.Article
	Tags                []*entity.Tag
	JsEnv               PageAdminArticleTagsJsEnv
}

type PageAdminArticleTagsJsEnv struct{}

func (t *Impl) PageAdminArticleTags(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	dto, err := t.U.GetAdminArticleTags(ctx, article)
	if err != nil {
		t.P.RenderHTML(
			ctx,
			http.StatusInternalServerError,
			"page_error.html",
			NewPageErrorUnknownError(),
		)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article_tags.html",
		PageAdminArticleTags{
			ComponentCommonHead: ComponentCommonHead{},
			JsEnv:               PageAdminArticleTagsJsEnv{},
			Article:             article,
			Tags:                dto.Tags,
		},
	)
}

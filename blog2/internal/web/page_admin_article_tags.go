package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageAdminArticleTags struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
	Article             *entity.Article
	Tags                []*entity.Tag
	JsEnv               PageAdminArticleTagsJsEnv
}

type PageAdminArticleTagsJsEnv struct{}

func (t *Impl) PageAdminArticleTags(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	dto, err := t.U.GetAdminArticleTags(ctx, article)
	if err != nil {
		t.L.Error("Failed to get admin article tags", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article_tags.html",
		PageAdminArticleTags{
			ComponentCommonHead: ComponentCommonHead{},
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			JsEnv:   PageAdminArticleTagsJsEnv{},
			Article: article,
			Tags:    dto.Tags,
		},
	)
}

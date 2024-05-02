package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageAdminArticle struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
	Article             *entity.Article
	JsEnv               PageAdminArticleJsEnv
}

type PageAdminArticleJsEnv struct {
	Article              entity.Article `json:"article"`
	NotAttachedTags      []*entity.Tag  `json:"notAttachedTags"`
	Markdown             string         `json:"markdown"`
	HTML                 string         `json:"html"`
	BaseURLFile          string         `json:"baseUrlFile"`
	BaseURLFileThumbnail string         `json:"baseUrlFileThumbnail"`
}

func (t *Impl) PageAdminArticle(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	dto, err := t.U.GetAdminArticle(ctx, article)
	if err != nil {
		t.L.Error("", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.RenderHTML(
		ctx,
		http.StatusOK,
		"page_admin_article.html",
		PageAdminArticle{
			ComponentHeader: ComponentHeader{
				IsAdmin: ctxGetAdmin(ctx),
			},
			ComponentCommonHead: ComponentCommonHead{},
			JsEnv: PageAdminArticleJsEnv{
				Article:              *article,
				NotAttachedTags:      dto.NotAttachedTags,
				Markdown:             dto.MarkdownBody,
				HTML:                 dto.HTMLBody,
				BaseURLFile:          t.BaseURLFile,
				BaseURLFileThumbnail: t.BaseURLFileThumbnail,
			},
		},
	)
}

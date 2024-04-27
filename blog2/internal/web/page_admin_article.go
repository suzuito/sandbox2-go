package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type PageAdminArticle struct {
	ComponentCommonHead ComponentCommonHead
	ComponentHeader     ComponentHeader
	Article             *entity.Article
	MarkdownBody        string
	HTMLBody            template.HTML
	JsEnv               PageAdminArticleJsEnv
}

type PageAdminArticleJsEnv struct {
	ArticleID entity.ArticleID `json:"articleId"`
	Published bool             `json:"published"`
}

func (t *Impl) PageAdminArticle(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	dto, err := t.U.GetAdminArticle(ctx, article.ID)
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
				ArticleID: article.ID,
				Published: article.Published,
			},
			Article:      article,
			MarkdownBody: dto.MarkdownBody,
			HTMLBody:     template.HTML(dto.HTMLBody),
		},
	)
}

func (t *Impl) PostAdminArticleEditTags(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	tagIDsAddString, _ := ctx.GetPostFormArray("add")
	tagIDsDeleteString, _ := ctx.GetPostFormArray("delete")
	tagIDsAdd := entity.NewTagIDs(tagIDsAddString)
	tagIDsDelete := entity.NewTagIDs(tagIDsDeleteString)
	if err := t.U.PostAdminArticleEditTags(ctx, article.ID, tagIDsAdd, tagIDsDelete); err != nil {
		t.L.Error("", "err", err)
		t.RenderUnknownError(ctx)
		return
	}
	t.P.Redirect(ctx, http.StatusFound, fmt.Sprintf("/admin/articles/%s/tags", article.ID))
}

func (t *Impl) PutAdminArticle(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	message := struct {
		Title *string `json:"title"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if _, err := t.U.PutAdminArticle(ctx, article.ID, message.Title); err != nil {
		t.L.Error("", "err", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (t *Impl) PutAdminArticleMarkdown(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	if err := t.U.PutAdminArticleMarkdown(ctx, article.ID, ctx.Request.Body); err != nil {
		t.L.Error("", "err", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (t *Impl) PostAdminArticlePublish(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	if err := t.U.PostAdminArticlePublish(ctx, article.ID); err != nil {
		t.L.Error("", "err", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

func (t *Impl) DeleteAdminArticlePublish(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	if err := t.U.DeleteAdminArticlePublish(ctx, article.ID); err != nil {
		t.L.Error("", "err", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusOK)
}

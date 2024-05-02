package web

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

func (t *Impl) APIPutAdminArticle(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	message := struct {
		Title     *string `json:"title"`
		Published *bool   `json:"published"`
	}{}
	if err := ctx.BindJSON(&message); err != nil {
		t.P.RenderJSON(ctx, http.StatusInternalServerError, APIErrorResponse{})
		return
	}
	dto, err := t.U.APIPutAdminArticle(ctx, article.ID, message.Title, message.Published)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.RenderJSON(ctx, http.StatusInternalServerError, APIErrorResponse{})
		return
	}
	t.P.RenderJSON(ctx, http.StatusOK, dto.Article)
}

func (t *Impl) APIPostAdminArticleEditTags(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	message := struct {
		Add    []entity.TagID `json:"add"`
		Delete []entity.TagID `json:"delete"`
	}{}
	if err := ctx.Bind(&message); err != nil {
		t.P.RenderJSON(ctx, http.StatusBadRequest, APIErrorResponse{})
		return
	}
	dto, err := t.U.APIPostAdminArticleEditTags(ctx, article, message.Add, message.Delete)
	if err != nil {
		t.L.Error("", "err", err)
		t.P.RenderJSON(ctx, http.StatusInternalServerError, APIErrorResponse{})
		return
	}
	t.P.RenderJSON(ctx, http.StatusOK, struct {
		Article         *entity.Article `json:"article"`
		NotAttachedTags []*entity.Tag   `json:"notAttachedTags"`
	}{
		Article:         dto.Article,
		NotAttachedTags: dto.NotAttachedTags,
	})
}

func (t *Impl) APIPutAdminArticleMarkdown(ctx *gin.Context) {
	article := ctxGetArticle(ctx)
	markdownBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		t.P.RenderJSON(ctx, http.StatusBadRequest, APIErrorResponse{})
		return
	}
	dto, err := t.U.APIPutAdminArticleMarkdown(
		ctx,
		article.ID,
		bytes.NewBufferString(string(markdownBody)),
	)
	if err != nil {
		t.P.RenderJSON(ctx, http.StatusInternalServerError, APIErrorResponse{})
		return
	}
	t.P.RenderJSON(ctx, http.StatusOK, struct {
		MarkdownBody string `json:"markdown"`
		HTMLBody     string `json:"html"`
	}{
		MarkdownBody: dto.MarkdownBody,
		HTMLBody:     dto.HTMLBody,
	})
}

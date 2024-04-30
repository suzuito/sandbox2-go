package usecase

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOAPIPutAdminArticle struct {
	Article *entity.Article
}

func (t *Impl) APIPutAdminArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	title *string,
	published *bool,
) (*DTOAPIPutAdminArticle, error) {
	article, err := t.S.PutArticle(ctx, articleID, title, published, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPutAdminArticle{
		Article: article,
	}, nil
}

type DTOAPIPostAdminArticleEditTags struct {
	Article         *entity.Article
	NotAttachedTags []*entity.Tag
}

func (t *Impl) APIPostAdminArticleEditTags(
	ctx context.Context,
	article *entity.Article,
	add []entity.TagID,
	delete []entity.TagID,
) (*DTOAPIPostAdminArticleEditTags, error) {
	article, err := t.S.UpdateArticleTags(
		ctx,
		article.ID,
		add,
		delete,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	tags, err := t.S.GetNotAttachedArticleTags(ctx, article)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostAdminArticleEditTags{
		Article:         article,
		NotAttachedTags: tags,
	}, nil
}

type DTOAPIPutAdminArticleMarkdown struct {
	MarkdownBody string
	HTMLBody     string
}

func (t *Impl) APIPutAdminArticleMarkdown(
	ctx context.Context,
	articleID entity.ArticleID,
	markdown io.Reader,
) (*DTOAPIPutAdminArticleMarkdown, error) {
	markdownBuffer, err := io.ReadAll(markdown)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	htmlBody, err := t.S.PutArticleMarkdown(ctx, articleID, string(markdownBuffer))
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPutAdminArticleMarkdown{
		MarkdownBody: string(markdownBuffer),
		HTMLBody:     htmlBody,
	}, nil
}

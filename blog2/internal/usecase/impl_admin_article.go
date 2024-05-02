package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticle struct {
	MarkdownBody    string
	HTMLBody        string
	NotAttachedTags []*entity.Tag
}

func (t *Impl) GetAdminArticle(
	ctx context.Context,
	article *entity.Article,
) (*DTOGetAdminArticle, error) {
	markdownBody, htmlBody, err := t.S.GetArticleBody(ctx, article.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	notAttachedTags, err := t.S.GetNotAttachedArticleTags(ctx, article)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticle{
		MarkdownBody:    markdownBody,
		HTMLBody:        htmlBody,
		NotAttachedTags: notAttachedTags,
	}, nil
}

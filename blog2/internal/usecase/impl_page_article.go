package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type PageArticle struct {
	MarkdownBody string
	HTMLBody     string
}

func (t *Impl) PageArticle(
	ctx context.Context,
	article *entity.Article,
) (*PageArticle, error) {
	markdownBody, htmlBody, err := t.S.GetArticleBody(ctx, article.ID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &PageArticle{
		MarkdownBody: markdownBody,
		HTMLBody:     htmlBody,
	}, nil
}

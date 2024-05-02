package service

import (
	"bytes"
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) GetArticleBody(
	ctx context.Context,
	articleID entity.ArticleID,
) (string, string, error) {
	markdownBodyBuffer := bytes.NewBufferString("")
	if err := t.StorageArticle.GetArticle(ctx, articleID, markdownBodyBuffer); err != nil {
		return "", "", terrors.Wrap(err)
	}
	htmlBody := ""
	if err := t.Markdown2HTML.Generate(ctx, markdownBodyBuffer.String(), &htmlBody); err != nil {
		return "", "", terrors.Wrap(err)
	}
	return markdownBodyBuffer.String(), htmlBody, nil
}

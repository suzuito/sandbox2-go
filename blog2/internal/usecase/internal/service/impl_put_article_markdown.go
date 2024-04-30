package service

import (
	"bytes"
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) PutArticleMarkdown(
	ctx context.Context,
	articleID entity.ArticleID,
	markdownBody string,
) (string, error) {
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleID)
	if err != nil {
		return "", terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		return "", terrors.Wrapf("Document %s is not found", articleID)
	}
	if err := t.StorageArticle.PutArticle(ctx, articleID, bytes.NewBufferString(markdownBody)); err != nil {
		return "", terrors.Wrap(err)
	}
	htmlBody := ""
	if err := t.Markdown2HTML.Generate(ctx, markdownBody, &htmlBody); err != nil {
		return "", terrors.Wrap(err)
	}
	return htmlBody, nil
}

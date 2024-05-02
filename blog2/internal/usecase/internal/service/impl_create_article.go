package service

import (
	"bytes"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) CreateArticle(
	ctx context.Context,
) (*entity.Article, error) {
	articleID := entity.ArticleID(uuid.New().String())
	if err := t.StorageArticle.PutArticle(ctx, articleID, bytes.NewBuffer([]byte{})); err != nil {
		return nil, terrors.Wrap(err)
	}
	article, err := t.RepositoryArticle.CreateArticle(ctx, articleID, time.Now())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	t.L.DebugContext(ctx, "Created article", "article", article)
	return article, nil
}

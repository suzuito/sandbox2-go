package usecase

import (
	"bytes"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticles struct {
	Articles   []*entity.Article
	NextOffset *int
	PrevOffset *int
}

func (t *Impl) GetAdminArticles(
	ctx context.Context,
	query *entity.ArticleSearchQuery,
) (*DTOGetAdminArticles, error) {
	articles, prevOffset, nextOffset, err := t.RepositoryArticle.SearchArticles(ctx, query)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticles{
		Articles:   articles,
		PrevOffset: prevOffset,
		NextOffset: nextOffset,
	}, nil
}

type DTOPostAdminArticles struct {
	Article *entity.Article
}

func (t *Impl) PostAdminArticles(
	ctx context.Context,
) (*DTOPostAdminArticles, error) {
	articleID := entity.ArticleID(uuid.New().String())
	if err := t.StorageArticle.PutArticle(ctx, articleID, bytes.NewBuffer([]byte{})); err != nil {
		return nil, terrors.Wrap(err)
	}
	article, err := t.RepositoryArticle.CreateArticle(ctx, articleID, time.Now())
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	t.L.DebugContext(ctx, "Created article", "article", article)
	return &DTOPostAdminArticles{
		Article: article,
	}, nil
}

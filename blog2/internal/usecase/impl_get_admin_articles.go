package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticles struct {
	Articles []*entity.Article
}

func (t *Impl) GetAdminArticles(
	ctx context.Context,
	query *entity.ArticleSearchQuery,
) (*DTOGetAdminArticles, error) {
	indices, err := t.RepositoryArticleIndex.Search(ctx, query)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	articleIDs := []entity.ArticleID{}
	for _, index := range indices {
		articleIDs = append(articleIDs, index.ArticleID)
	}
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleIDs...)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticles{
		Articles: articles,
	}, nil
}

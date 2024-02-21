package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticle struct {
	Article *entity.Article
}

func (t *Impl) GetAdminArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*DTOGetAdminArticle, error) {
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		// TODO 404 error
		return nil, terrors.Wrapf("Not found")
	}
	return &DTOGetAdminArticle{
		Article: articles[0],
	}, nil
}

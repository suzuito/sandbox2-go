package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOMiddlewareGetArticle struct {
	Article *entity.Article
}

func (t *Impl) MiddlewareGetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*DTOMiddlewareGetArticle, error) {
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		return nil, &NotFoundEntityError{
			EntityType: "Article",
			EntityID:   string(articleID),
		}
	}
	return &DTOMiddlewareGetArticle{
		Article: articles[0],
	}, nil
}

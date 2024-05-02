package service

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase/pkg/serviceerror"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) GetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*entity.Article, error) {
	articles, err := t.RepositoryArticle.GetArticles(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		return nil, &serviceerror.NotFoundEntityError{
			EntityType: "Article",
			EntityID:   string(articleID),
		}
	}
	return articles[0], nil
}

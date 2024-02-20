package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryArticle interface {
	GetArticles(ctx context.Context, ids ...entity.ArticleID) ([]*entity.Article, error)
}

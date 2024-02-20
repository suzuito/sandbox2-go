package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type Usecase interface {
	GetAdminArticles(
		ctx context.Context,
		query *entity.ArticleSearchQuery,
	) (*DTOGetAdminArticles, error)
	GetAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*DTOGetAdminArticle, error)
}

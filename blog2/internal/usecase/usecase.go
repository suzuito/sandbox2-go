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
	PostAdminArticles(
		ctx context.Context,
	) (*DTOPostAdminArticles, error)
	GetAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*DTOGetAdminArticle, error)
	MiddlewareGetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*DTOMiddlewareGetArticle, error)
}

package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type RepositoryArticleSource interface {
	GetArticleSources(
		ctx context.Context,
		ref string,
		proc func(*entity.ArticleSource, []byte) error,
	) error
	SearchArticleSources(
		ctx context.Context,
		queryString string,
		proc func(*entity.ArticleSource) error,
	) error
	GetArticleSource(
		ctx context.Context,
		articleSourceID entity.ArticleSourceID,
		version string,
	) (*entity.ArticleSource, []byte, error)
	GetBranches(
		ctx context.Context,
	) ([]string, error)
	GetVersions(
		ctx context.Context,
		branch string,
		articleSourceID entity.ArticleSourceID,
	) ([]entity.ArticleSource, error)
}

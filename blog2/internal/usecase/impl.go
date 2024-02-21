package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/blog2/internal/repository"
)

type Impl struct {
	RepositoryArticleIndex repository.RepositoryArticleIndex
	RepositoryArticle      repository.RepositoryArticle
	L                      *slog.Logger
}

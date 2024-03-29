package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
	"github.com/suzuito/sandbox2-go/blog2/internal/repository"
)

type Impl struct {
	RepositoryArticle                  repository.RepositoryArticle
	StorageArticle                     repository.StorageArticle
	StorageArticleFileDirectlyUploaded repository.StorageArticleFileDirectlyUploaded
	Markdown2HTML                      markdown2html.Markdown2HTML
	L                                  *slog.Logger
}

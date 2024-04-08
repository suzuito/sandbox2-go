package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/imageconverter"
	"github.com/suzuito/sandbox2-go/blog2/internal/queue"
	"github.com/suzuito/sandbox2-go/blog2/internal/repository"
)

type Impl struct {
	RepositoryArticle                       repository.RepositoryArticle
	StorageArticle                          repository.StorageArticle
	StorageArticleFileUploaded              repository.StorageArticleFileUploaded
	StorageArticleFile                      repository.StorageArticleFile
	RepositoryArticleFileUploaded           repository.RepositoryArticleFileUploaded
	ArticleFileImageConverter               imageconverter.ImageConverter
	FunctionTriggerStartFileUploadedProcess queue.FunctionTrigger
	Markdown2HTML                           markdown2html.Markdown2HTML
	L                                       *slog.Logger
}

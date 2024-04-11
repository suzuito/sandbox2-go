package service

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/imageconverter"
	"github.com/suzuito/sandbox2-go/blog2/internal/repository"
)

type Impl struct {
	RepositoryArticle      repository.RepositoryArticle
	StorageArticle         repository.StorageArticle
	StorageFileUploaded    repository.StorageFileUploaded
	StorageFile            repository.StorageFile
	RepositoryFileUploaded repository.RepositoryFileUploaded
	FileImageConverter     imageconverter.ImageConverter
	Markdown2HTML          markdown2html.Markdown2HTML
	L                      *slog.Logger
}

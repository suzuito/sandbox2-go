package usecase

import (
	"log/slog"

	"github.com/suzuito/sandbox2-go/blog2/internal/markdown2html"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/filetypedetector"
	"github.com/suzuito/sandbox2-go/blog2/internal/procs/articlefile/pkg/imageconverter"
	"github.com/suzuito/sandbox2-go/blog2/internal/repository"
	internal_service "github.com/suzuito/sandbox2-go/blog2/internal/usecase/internal/service"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase/pkg/service"
)

type Impl struct {
	S service.Service
	L *slog.Logger
}

func NewImpl(
	repositoryArticle repository.RepositoryArticle,
	storageArticle repository.StorageArticle,
	storageFileUploaded repository.StorageFileUploaded,
	storageFile repository.StorageFile,
	storageFileThumbnail repository.StorageFileThumbnal,
	repositoryFileUploaded repository.RepositoryFileUploaded,
	fileImageConverter imageconverter.ImageConverter,
	markdown2HTML markdown2html.Markdown2HTML,
	fileTypeDetector filetypedetector.FileTypeDetector,
	logger *slog.Logger,
) *Impl {
	return &Impl{
		S: &internal_service.Impl{
			RepositoryArticle:      repositoryArticle,
			StorageArticle:         storageArticle,
			StorageFileUploaded:    storageFileUploaded,
			StorageFile:            storageFile,
			StorageFileThumbnail:   storageFileThumbnail,
			RepositoryFileUploaded: repositoryFileUploaded,
			FileImageConverter:     fileImageConverter,
			Markdown2HTML:          markdown2HTML,
			FileTypeDetector:       fileTypeDetector,
			L:                      logger,
		},
		L: logger,
	}
}

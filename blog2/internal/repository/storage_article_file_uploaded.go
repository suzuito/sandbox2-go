package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticleFileUploaded interface {
	Put(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, r io.Reader) error
	Get(
		ctx context.Context,
		articleID entity.ArticleID,
		fileID entity.ArticleFileUploadedID,
		w io.Writer,
	) error
	GetReader(
		ctx context.Context,
		articleID entity.ArticleID,
		fileID entity.ArticleFileUploadedID,
		f func(r io.Reader) error,
	) error
}

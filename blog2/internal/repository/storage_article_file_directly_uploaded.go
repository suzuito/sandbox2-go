package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticleFileDirectlyUploaded interface {
	Put(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileDirectlyUploadedID, r io.Reader) error
	Get(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileDirectlyUploadedID, w io.Writer) error
}

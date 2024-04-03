package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticleFile interface {
	Put(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, r io.Reader) error
	Get(ctx context.Context, articleID entity.ArticleID, fileID entity.ArticleFileUploadedID, w io.Writer) error
}

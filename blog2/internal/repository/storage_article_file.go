package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticleFile interface {
	Put(
		ctx context.Context,
		articleID entity.ArticleID,
		file *entity.ArticleFile,
		r io.Reader,
	) error
	PutThumbnail(
		ctx context.Context,
		articleID entity.ArticleID,
		file *entity.ArticleFileThumbnail,
		r io.Reader,
	) error
}

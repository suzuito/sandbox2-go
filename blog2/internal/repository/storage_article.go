package repository

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticle interface {
	PutArticle(ctx context.Context, articleID entity.ArticleID, r io.Reader) error
	GetArticle(ctx context.Context, articleID entity.ArticleID, w io.Writer) error
}

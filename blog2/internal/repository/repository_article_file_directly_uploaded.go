package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryArticleFileDirectlyUploaded interface {
	Put(ctx context.Context, file *entity.ArticleFileDirectlyUploaded) error
}

package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryArticleFileUploaded interface {
	Create(ctx context.Context, f *entity.ArticleFileUploaded) error
	Get(
		ctx context.Context,
		articleID entity.ArticleID,
		fileID entity.ArticleFileUploadedID,
	) (*entity.ArticleFileUploaded, error)
}

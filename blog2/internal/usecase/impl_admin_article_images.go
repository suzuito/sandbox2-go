package usecase

import (
	"context"
	"io"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOPostAdminArticleImages struct {
	Image entity.ArticleImage
}

func (t *Impl) PostAdminArticleImages(
	ctx context.Context,
	article *entity.Article,
	input io.Reader,
) (*DTOPostAdminArticleImages, error) {
	file := entity.ArticleFileDirectlyUploaded{
		ID: entity.ArticleFileDirectlyUploadedID(uuid.New().String()),
	}
	if err := t.StorageArticleFileDirectlyUploaded.Put(ctx, article.ID, file.ID, input); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPostAdminArticleImages{}, nil
}

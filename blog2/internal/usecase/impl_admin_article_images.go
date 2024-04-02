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
	file := entity.ArticleFileUploaded{
		ArticleID: article.ID,
		ID:        entity.ArticleFileUploadedID(uuid.New().String()),
	}
	// TODO Impl transaction
	if err := t.RepositoryArticleFileUploaded.Create(ctx, &file); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.StorageArticleFileUploaded.Put(ctx, article.ID, file.ID, input); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.FunctionTriggerStartImageProcess.Put(ctx, &file); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPostAdminArticleImages{}, nil
}

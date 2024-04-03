package infra

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type RepositoryArticleFileUploaded struct {
	Cli *firestore.Client
}

func (t *RepositoryArticleFileUploaded) Create(
	ctx context.Context,
	file *entity.ArticleFileUploaded,
) error {
	return t.Cli.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		dr := t.Cli.Collection("Blog2Articles").Doc(string(file.ArticleID)).Collection("FilesUploaded").Doc(string(file.ID))
		if err := tx.Set(dr, file); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	})
}

func (t *RepositoryArticleFileUploaded) Get(
	ctx context.Context,
	articleID entity.ArticleID,
	fileID entity.ArticleFileUploadedID,
) (*entity.ArticleFileUploaded, error) {
	dr := t.Cli.Collection("Blog2Articles").Doc(string(articleID)).Collection("FilesUploaded").Doc(string(fileID))
	snp, err := dr.Get(ctx)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	file := entity.ArticleFileUploaded{}
	if err := snp.DataTo(&file); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}

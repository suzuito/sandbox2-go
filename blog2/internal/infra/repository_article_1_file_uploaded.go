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

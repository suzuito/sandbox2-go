package infra

import (
	"context"

	"github.com/google/uuid"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) CreateTag(ctx context.Context, name string) (*entity.Tag, error) {
	tag := entity.Tag{}
	tagID := uuid.New().String()
	err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		_, err := csql.ExecContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"INSERT IGNORE INTO `tags`(`id`, `name`, `created_at`, `updated_at`) VALUES (?, ?, NOW(), NOW())",
			tagID,
			name,
		)
		if err != nil {
			return terrors.Wrap(err)
		}
		rows, err := csql.QueryContext(
			ctx, tx,
			"SELECT `id`, `name` FROM `tags` WHERE `name` = ?",
			name,
		)
		if err != nil {
			return terrors.Wrap(err)
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
				return terrors.Wrap(err)
			}
			break
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &tag, nil
}

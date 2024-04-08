package infra

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) PutFile(
	ctx context.Context,
	file *entity.File,
) error {
	if err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		if _, err := execContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"INSERT IGNORE INTO `files`(`id`, `name`, `type`, `media_type`, `exists_thumbnail`, `created_at`)"+
				"VALUES (?, ?, ?, ?, ?, NOW())",
			file.ID,
			file.Name,
			file.Type,
			file.MediaType,
			file.ExistsThumbnail,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *RepositoryArticle) DeleteFile(
	ctx context.Context,
	fileID entity.FileID,
) error {
	if err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		if _, err := execContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"DELETE FROM `files` WHERE `id` = ?",
			fileID,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *RepositoryArticle) GetArticleFiles(
	ctx context.Context,
	articleID entity.ArticleID,
) error {
	return terrors.Wrapf("not impl")
}

package infra

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) PutFileThumbnail(
	ctx context.Context,
	file *entity.FileThumbnail,
) error {
	if err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		if _, err := execContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"INSERT IGNORE INTO `file_thumbnails`(`id`, `file_id`, `media_type`)"+
				"VALUES (?, ?, ?)",
			file.ID,
			file.FileID,
			file.MediaType,
		); err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *RepositoryArticle) DeleteFileThumbnail(
	ctx context.Context,
	fileID entity.FileID,
) error {
	if err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		if _, err := execContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"DELETE FROM `file_thumbnails` WHERE `file_id` = ?",
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

func (t *RepositoryArticle) GetFileThumbnail(
	ctx context.Context,
	fileID entity.FileID,
) (*entity.FileThumbnail, error) {
	// Get tags
	row := queryRowContext(
		ctx,
		t.Pool,
		"SELECT `id`,`file_id`,`media_type` FROM `file_thumbnails` WHERE id = ?",
		fileID,
	)
	file := entity.FileThumbnail{}
	if err := row.Scan(
		&file.ID,
		&file.FileID,
		&file.MediaType,
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := row.Err(); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}

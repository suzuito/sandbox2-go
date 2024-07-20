package infra

import (
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) PutFile(
	ctx context.Context,
	file *entity.File,
) error {
	if err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		if _, err := csql.ExecContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"INSERT IGNORE INTO `files`(`id`, `type`, `media_type`, `created_at`, `updated_at`)"+
				"VALUES (?, ?, ?, NOW(), NOW())",
			file.ID,
			file.Type,
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

func (t *RepositoryArticle) DeleteFile(
	ctx context.Context,
	fileID entity.FileID,
) error {
	if err := csql.WithTransaction(ctx, t.Pool, func(tx csql.TxOrDB) error {
		if _, err := csql.ExecContext(
			ctx,
			tx,
			// 既存のタグがある場合IGNOREする
			"DELETE FROM `file_thumbnails` WHERE `file_id` = ?",
			fileID,
		); err != nil {
			return terrors.Wrap(err)
		}
		if _, err := csql.ExecContext(
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

func (t *RepositoryArticle) GetFile(
	ctx context.Context,
	fileID entity.FileID,
) (*entity.File, error) {
	// Get tags
	row := csql.QueryRowContext(
		ctx,
		t.Pool,
		"SELECT `id`,`type`,`media_type` FROM `files` WHERE id = ?",
		fileID,
	)
	file := entity.File{}
	if err := row.Scan(
		&file.ID,
		&file.Type,
		&file.MediaType,
	); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := row.Err(); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}

func (t *RepositoryArticle) SearchFiles(
	ctx context.Context,
	queryString string,
	offset int,
	limit int,
) ([]*entity.FileAndThumbnail, error) {
	fmt.Printf("=%s=%d=%d\n", queryString, offset, limit)
	args := []any{}
	query := "SELECT `id`,`type`,`media_type`,`created_at`,`updated_at` FROM `files`"
	if queryString != "" {
		args = append(args, queryString)
		query += " WHERE MATCH (id) AGAINST (? IN BOOLEAN MODE)"
	}
	query += " ORDER BY `updated_at` DESC"
	args = append(args, limit, offset)
	query += " LIMIT ? OFFSET ?"
	// Get files
	rowsFiles, err := csql.QueryContext(
		ctx,
		t.Pool,
		query,
		args...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rowsFiles.Close()
	files := []*entity.File{}
	fileIDs := []string{}
	for rowsFiles.Next() {
		file := entity.File{}
		if err := rowsFiles.Scan(
			&file.ID,
			&file.Type,
			&file.MediaType,
			&file.CreatedAt,
			&file.UpdatedAt,
		); err != nil {
			return nil, terrors.Wrap(err)
		}
		files = append(files, &file)
		fileIDs = append(fileIDs, string(file.ID))
	}
	// Get thumbnails
	thumbnailMap := map[entity.FileID]*entity.FileThumbnail{}
	if len(fileIDs) > 0 {
		rowsThumbnails, err := csql.QueryContext(
			ctx,
			t.Pool,
			fmt.Sprintf(
				"SELECT `id`,`file_id`,`media_type` FROM `file_thumbnails` WHERE %s",
				csql.SqlIn("file_id", fileIDs),
			),
			csql.ToAnySlice(fileIDs)...,
		)
		if err != nil {
			return nil, terrors.Wrap(err)
		}
		defer rowsThumbnails.Close()
		for rowsThumbnails.Next() {
			fileThumbnail := entity.FileThumbnail{}
			if err := rowsThumbnails.Scan(
				&fileThumbnail.ID,
				&fileThumbnail.FileID,
				&fileThumbnail.MediaType,
			); err != nil {
				return nil, terrors.Wrap(err)
			}
			thumbnailMap[fileThumbnail.FileID] = &fileThumbnail
		}
	}

	fileAndThumbnails := []*entity.FileAndThumbnail{}
	for _, file := range files {
		thumbnail, exists := thumbnailMap[file.ID]
		fileAndThumbnail := entity.FileAndThumbnail{
			File: file,
		}
		if exists {
			fileAndThumbnail.FileThumbnail = thumbnail
		}
		fileAndThumbnails = append(fileAndThumbnails, &fileAndThumbnail)
	}
	return fileAndThumbnails, nil
}

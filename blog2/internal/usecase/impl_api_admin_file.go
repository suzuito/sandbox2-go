package usecase

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOAPIPostAdminFiles struct {
	File          *entity.File
	FileThumbnail *entity.FileThumbnail
}

func (t *Impl) APIPostAdminFiles(
	ctx context.Context,
	fileName string,
	file io.Reader,
) (*DTOAPIPostAdminFiles, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		t.L.Error("", "err", err)
		return nil, terrors.Wrap(err)
	}
	ft, mediaType := t.S.DetectFileType(ctx, data)
	articleFile := entity.File{
		ID:        entity.FileID(fileName),
		Type:      ft,
		MediaType: mediaType,
	}
	// Save file
	if err := t.S.ExistFile(
		ctx,
		articleFile.ID,
	); err != nil {
		t.L.Error("", "err", err)
		return nil, terrors.Wrap(err)
	}
	if err := t.S.PutFile(
		ctx,
		&articleFile,
		data,
	); err != nil {
		t.L.Error("", "err", err)
		return nil, terrors.Wrap(err)
	}
	// Create thumbnail
	fileThumbnail, err := t.S.CreateThumbnail(ctx, articleFile.ID)
	if err != nil {
		t.L.Error("", "err", err)
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIPostAdminFiles{
		File:          &articleFile,
		FileThumbnail: fileThumbnail,
	}, nil
}

type DTOAPIGetAdminFiles struct {
	Files []*entity.FileAndThumbnail
}

func (t *Impl) APIGetAdminFiles(
	ctx context.Context,
	queryString string,
	page int,
	limit int,
) (*DTOAPIGetAdminFiles, error) {
	files, err := t.S.SearchFiles(ctx, queryString, page*limit, limit)
	if err != nil {
		t.L.Error("", "err", err)
		return nil, terrors.Wrap(err)
	}
	return &DTOAPIGetAdminFiles{
		Files: files,
	}, nil
}

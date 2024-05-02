package service

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"mime"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase/pkg/serviceerror"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) DetectFileType(
	ctx context.Context,
	data []byte,
) (entity.FileType, string) {
	return t.FileTypeDetector.Do(data)
}

func (t *Impl) ExistFile(
	ctx context.Context,
	fileID entity.FileID,
) error {
	_, err := t.RepositoryArticle.GetFile(ctx, fileID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return terrors.Wrap(err)
	}
	return terrors.Wrap(&serviceerror.AlreadyExistsEntityError{
		EntityType: "file",
		EntityID:   string(fileID),
	})
}

func (t *Impl) PutFile(
	ctx context.Context,
	file *entity.File,
	data []byte,
) error {
	// TODO トランザクション使って中途半端な状態になることを防止すべき箇所
	if err := terrors.Wrap(t.StorageFile.Put(ctx, file, bytes.NewBuffer(data))); err != nil {
		return terrors.Wrap(err)
	}
	if err := t.RepositoryArticle.PutFile(ctx, file); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) CreateThumbnail(
	ctx context.Context,
	fileID entity.FileID,
) (*entity.FileThumbnail, error) {
	file, err := t.RepositoryArticle.GetFile(ctx, fileID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	var fileThumbnail *entity.FileThumbnail
	var dataThumbnail []byte
	existsThumbnail := false
	switch file.Type {
	case entity.FileTypeImage:
		fileThumbnail, dataThumbnail, err = t.createThumbnailImage(ctx, file)
		existsThumbnail = true
	}
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if !existsThumbnail {
		t.L.Debug("Thumbnail is not created", "fileID", fileID)
		return nil, nil
	}
	// TODO トランザクション使うべき
	if err := t.StorageFileThumbnail.Put(ctx, fileThumbnail, bytes.NewBuffer(dataThumbnail)); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.RepositoryArticle.PutFileThumbnail(ctx, fileThumbnail); err != nil {
		return nil, terrors.Wrap(err)
	}
	return fileThumbnail, nil
}

func (t *Impl) createThumbnailImage(
	ctx context.Context,
	file *entity.File,
) (*entity.FileThumbnail, []byte, error) {
	data := []byte{}
	dataBuffer := bytes.NewBuffer(data)
	if err := t.StorageFile.Get(ctx, file, dataBuffer); err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	img, _, thumbnailEncoder, err := t.FileImageConverter.Decode(dataBuffer)
	if err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	thumbnail := t.FileImageConverter.CreateThumbnail(img)
	mediaType := thumbnailEncoder.GetMediaType()
	ext := ""
	exts, err := mime.ExtensionsByType(mediaType)
	if err == nil && len(exts) > 0 {
		ext = exts[0]
	} else {
		t.L.Debug("mime type is unknown", "mediaType", mediaType)
	}
	fileThumbnail := entity.FileThumbnail{
		ID:        entity.FileThumbnailID(string(file.ID) + ext),
		FileID:    file.ID,
		MediaType: mediaType,
	}
	thumbnailBuffer := bytes.NewBuffer([]byte{})
	if err := thumbnailEncoder.Encode(thumbnailBuffer, thumbnail); err != nil {
		return nil, nil, terrors.Wrap(err)
	}
	return &fileThumbnail, thumbnailBuffer.Bytes(), nil
}

func (t *Impl) SearchFiles(
	ctx context.Context,
	queryString string,
	offset int,
	limit int,
) ([]*entity.FileAndThumbnail, error) {
	return t.RepositoryArticle.SearchFiles(ctx, queryString, offset, limit)
}

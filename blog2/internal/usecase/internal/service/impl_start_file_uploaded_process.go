package service

import (
	"bytes"
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) StartFileUploadedProcess(ctx context.Context, fileUploaded *entity.FileUploaded) (*entity.File, error) {
	t.L.Info("Hello world!", "file", fileUploaded)
	var file *entity.File
	_, err := t.RepositoryFileUploaded.StartProcess(ctx, fileUploaded.ID, func() error {
		var err error
		switch fileUploaded.Type {
		case entity.FileTypeImage:
			file, err = t.StartFileUploadedImageProcess(ctx, fileUploaded)
		default:
			file, err = t.StartFileUploadedUnknownProcess(ctx, fileUploaded)
		}
		return err
	})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return file, nil
}

func (t *Impl) StartFileUploadedImageProcess(ctx context.Context, fileUploaded *entity.FileUploaded) (*entity.File, error) {
	if fileUploaded.Type != entity.FileTypeImage {
		return nil, terrors.Wrapf("invalid file type %s", fileUploaded.Type)
	}
	srcImageBytes := bytes.NewBuffer([]byte{})
	if err := t.StorageFileUploaded.Get(ctx, fileUploaded.ID, srcImageBytes); err != nil {
		return nil, terrors.Wrap(err)
	}
	// ここの処理、どう書くのが良いか？いまいちよくわからん。
	// なのでベタ書き。後でリファクタリングする。
	dstImage, dstEncoder, thumbnailEncoder, err := t.FileImageConverter.Decode(srcImageBytes)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	file := entity.File{
		ID:              entity.FileID(fileUploaded.ID),
		Name:            fileUploaded.Name,
		Type:            fileUploaded.Type,
		MediaType:       dstEncoder.GetMediaType(),
		ExistsThumbnail: true,
	}
	// 画像変換
	// Small画像の生成
	dstImageBytes := bytes.NewBuffer([]byte{})
	if err := dstEncoder.Encode(dstImageBytes, dstImage); err != nil {
		return nil, terrors.Wrap(err)
	}
	if err := t.StorageFile.Put(ctx, &file, dstImageBytes); err != nil {
		return nil, terrors.Wrap(err)
	}
	dstImageThumbnail := t.FileImageConverter.CreateThumbnail(dstImage)
	dstImageThumbnailBytes := bytes.NewBuffer([]byte{})
	if err := thumbnailEncoder.Encode(dstImageThumbnailBytes, dstImageThumbnail); err != nil {
		return nil, terrors.Wrap(err)
	}
	articleFileThumbnail := entity.FileThumbnail{
		ID:        file.ID,
		MediaType: thumbnailEncoder.GetMediaType(),
	}
	if err := t.StorageFile.PutThumbnail(ctx, &articleFileThumbnail, dstImageThumbnailBytes); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}

func (t *Impl) StartFileUploadedUnknownProcess(ctx context.Context, fileUploaded *entity.FileUploaded) (*entity.File, error) {
	fileData := []byte{}
	fileDataBuffer := bytes.NewBuffer(fileData)
	if err := t.StorageFileUploaded.Get(ctx, fileUploaded.ID, fileDataBuffer); err != nil {
		return nil, terrors.Wrap(err)
	}
	file := entity.File{
		ID:              entity.FileID(fileUploaded.ID),
		Name:            fileUploaded.Name,
		Type:            fileUploaded.Type,
		MediaType:       "",
		ExistsThumbnail: false,
	}
	if err := t.StorageFile.Put(ctx, &file, fileDataBuffer); err != nil {
		return nil, terrors.Wrap(err)
	}
	return &file, nil
}

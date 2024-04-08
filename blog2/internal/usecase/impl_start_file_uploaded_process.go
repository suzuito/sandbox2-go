package usecase

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) StartFileUploadedProcessFromGCF(ctx context.Context, data []byte) error {
	file := entity.FileUploaded{}
	if err := json.Unmarshal(data, &file); err != nil {
		return terrors.Wrap(err)
	}
	return t.StartFileUploadedProcess(ctx, &file)
}

func (t *Impl) StartFileUploadedProcess(ctx context.Context, file *entity.FileUploaded) error {
	t.L.Info("Hello world!", "file", file)
	_, err := t.RepositoryFileUploaded.StartProcess(ctx, file.ID, func() error {
		switch file.Type {
		case entity.FileTypeImage:
			return t.StartFileUploadedImageProcess(ctx, file)
		}
		return t.StartFileUploadedUnknownProcess(ctx, file)
	})
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) StartFileUploadedImageProcess(ctx context.Context, fileUploaded *entity.FileUploaded) error {
	if fileUploaded.Type != entity.FileTypeImage {
		return terrors.Wrapf("invalid file type %s", fileUploaded.Type)
	}
	srcImageBytes := bytes.NewBuffer([]byte{})
	if err := t.StorageFileUploaded.Get(ctx, fileUploaded.ID, srcImageBytes); err != nil {
		return terrors.Wrap(err)
	}
	// ここの処理、どう書くのが良いか？いまいちよくわからん。
	// なのでベタ書き。後でリファクタリングする。
	dstImage, dstEncoder, thumbnailEncoder, err := t.FileImageConverter.Decode(srcImageBytes)
	if err != nil {
		return terrors.Wrap(err)
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
		return terrors.Wrap(err)
	}
	if err := t.StorageFile.Put(ctx, &file, dstImageBytes); err != nil {
		return terrors.Wrap(err)
	}
	dstImageThumbnail := t.FileImageConverter.CreateThumbnail(dstImage)
	dstImageThumbnailBytes := bytes.NewBuffer([]byte{})
	if err := thumbnailEncoder.Encode(dstImageThumbnailBytes, dstImageThumbnail); err != nil {
		return terrors.Wrap(err)
	}
	articleFileThumbnail := entity.FileThumbnail{
		ID:        file.ID,
		MediaType: thumbnailEncoder.GetMediaType(),
	}
	if err := t.StorageFile.PutThumbnail(ctx, &articleFileThumbnail, dstImageThumbnailBytes); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) StartFileUploadedUnknownProcess(ctx context.Context, fileUploaded *entity.FileUploaded) error {
	fileData := []byte{}
	fileDataBuffer := bytes.NewBuffer(fileData)
	if err := t.StorageFileUploaded.Get(ctx, fileUploaded.ID, fileDataBuffer); err != nil {
		return terrors.Wrap(err)
	}
	file := entity.File{
		ID:              entity.FileID(fileUploaded.ID),
		Name:            fileUploaded.Name,
		Type:            fileUploaded.Type,
		MediaType:       "",
		ExistsThumbnail: false,
	}
	if err := t.StorageFile.Put(ctx, &file, fileDataBuffer); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

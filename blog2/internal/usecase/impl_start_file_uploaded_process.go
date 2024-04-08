package usecase

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) StartFileUploadedProcessFromGCF(ctx context.Context, data []byte) error {
	file := entity.ArticleFileUploaded{}
	if err := json.Unmarshal(data, &file); err != nil {
		return terrors.Wrap(err)
	}
	return t.StartFileUploadedProcess(ctx, &file)
}

func (t *Impl) StartFileUploadedProcess(ctx context.Context, file *entity.ArticleFileUploaded) error {
	t.L.Info("Hello world!", "file", file)
	switch file.Type {
	case entity.ArticleFileTypeImage:
		return t.StartFileUploadedImageProcess(ctx, file)
	}
	return t.StartFileUploadedUnknownProcess(ctx, file)
}

func (t *Impl) StartFileUploadedImageProcess(ctx context.Context, file *entity.ArticleFileUploaded) error {
	if file.Type != entity.ArticleFileTypeImage {
		return terrors.Wrapf("invalid file type %s", file.Type)
	}
	srcImageBytes := bytes.NewBuffer([]byte{})
	if err := t.StorageArticleFileUploaded.Get(ctx, file.ArticleID, file.ID, srcImageBytes); err != nil {
		return terrors.Wrap(err)
	}
	// ここの処理、どう書くのが良いか？いまいちよくわからん。
	// なのでベタ書き。後でリファクタリングする。
	dstImage, dstEncoder, thumbnailEncoder, err := t.ArticleFileImageConverter.Decode(srcImageBytes)
	if err != nil {
		return terrors.Wrap(err)
	}
	articleFile := entity.ArticleFile{
		ID:        entity.ArticleFileID(file.ID),
		Type:      file.Type,
		MediaType: dstEncoder.GetMediaType(),
	}
	// 画像変換
	// Small画像の生成
	dstImageBytes := bytes.NewBuffer([]byte{})
	if err := dstEncoder.Encode(dstImageBytes, dstImage); err != nil {
		return terrors.Wrap(err)
	}
	if err := t.StorageArticleFile.Put(ctx, file.ArticleID, &articleFile, dstImageBytes); err != nil {
		return terrors.Wrap(err)
	}
	dstImageThumbnail := t.ArticleFileImageConverter.CreateThumbnail(dstImage)
	dstImageThumbnailBytes := bytes.NewBuffer([]byte{})
	if err := thumbnailEncoder.Encode(dstImageBytes, dstImageThumbnail); err != nil {
		return terrors.Wrap(err)
	}
	articleFileThumbnail := entity.ArticleFileThumbnail{
		ID:        articleFile.ID,
		MediaType: thumbnailEncoder.GetMediaType(),
	}
	if err := t.StorageArticleFile.PutThumbnail(ctx, file.ArticleID, &articleFileThumbnail, dstImageThumbnailBytes); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

func (t *Impl) StartFileUploadedUnknownProcess(ctx context.Context, file *entity.ArticleFileUploaded) error {
	fileData := []byte{}
	fileDataBuffer := bytes.NewBuffer(fileData)
	if err := t.StorageArticleFileUploaded.Get(ctx, file.ArticleID, file.ID, fileDataBuffer); err != nil {
		return terrors.Wrap(err)
	}
	articleFile := entity.ArticleFile{
		ID:        entity.ArticleFileID(file.ID),
		Type:      file.Type,
		MediaType: "",
	}
	if err := t.StorageArticleFile.Put(ctx, file.ArticleID, &articleFile, fileDataBuffer); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

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
	fileData := []byte{}
	fileDataBuffer := bytes.NewBuffer(fileData)
	var formatName string
	if err := t.StorageArticleFileUploaded.GetReader(ctx, file.ArticleID, file.ID, func(r io.Reader) error {
		var err error
		_, formatName, err = image.Decode(fileDataBuffer)
		if err != nil {
			return terrors.Wrap(err)
		}
		return nil
	}); err != nil {
		return terrors.Wrap(err)
	}
	// TODO 画像を加工する場合、ここでやる
	mediaType := ""
	switch formatName {
	case "jpeg": // Format name https://github.com/golang/go/blob/db6097f8cbaceaed02051850d2411c88b763a0c3/src/image/jpeg/reader.go#L811
		mediaType = "image/jpeg"
	}
	articleFile := entity.ArticleFile{
		ID:        entity.ArticleFileID(file.ID),
		Type:      file.Type,
		MediaType: mediaType,
	}
	if err := t.StorageArticleFile.Put(ctx, file.ArticleID, &articleFile, fileDataBuffer); err != nil {
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

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
	t.L.Info("Hello world!", "file", file)
	fileData := []byte{}
	fileDataBuffer := bytes.NewBuffer(fileData)
	if err := t.StorageArticleFileUploaded.Get(ctx, file.ArticleID, file.ID, fileDataBuffer); err != nil {
		return terrors.Wrap(err)
	}
	if err := t.StorageArticleFile.Put(ctx, file.ArticleID, file.ID, fileDataBuffer); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

package usecase

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *Impl) StartImageProcessFromGCF(ctx context.Context, data []byte) error {
	file := entity.ArticleFileUploaded{}
	if err := json.Unmarshal(data, &file); err != nil {
		return terrors.Wrap(err)
	}
	t.L.Info("Hello world!", "file", file)
	imageBytes := []byte{}
	imageBytesBuffer := bytes.NewBuffer(imageBytes)
	if err := t.StorageArticleFileUploaded.Get(ctx, file.ArticleID, file.ID, imageBytesBuffer); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

package infra

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type StorageArticleFile struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageArticleFile) filePath(articleID entity.ArticleID, file *entity.ArticleFile) string {
	return fmt.Sprintf("%s/%s%s", articleID, file.ID, file.ExtIncludingDot())
}

func (t *StorageArticleFile) Put(
	ctx context.Context,
	articleID entity.ArticleID,
	file *entity.ArticleFile,
	r io.Reader,
) error {
	w := t.Cli.Bucket(t.Bucket).Object(t.filePath(articleID, file)).NewWriter(ctx)
	if file.MediaType != "" {
		w.ContentType = file.MediaType
	}
	_, err := io.Copy(w, r)
	if err != nil {
		return terrors.Wrap(err)
	}
	if err := w.Close(); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

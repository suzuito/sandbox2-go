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

func (t *StorageArticleFile) filePathThumbnail(articleID entity.ArticleID, file *entity.ArticleFileThumbnail) string {
	return fmt.Sprintf("%s/%s_thumbnail%s", articleID, file.ID, file.ExtIncludingDot())
}

func (t *StorageArticleFile) put(
	ctx context.Context,
	filePath string,
	mediaType string,
	r io.Reader,
) error {
	w := t.Cli.Bucket(t.Bucket).Object(filePath).NewWriter(ctx)
	if mediaType != "" {
		w.ContentType = mediaType
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

func (t *StorageArticleFile) Put(
	ctx context.Context,
	articleID entity.ArticleID,
	file *entity.ArticleFile,
	r io.Reader,
) error {
	return terrors.Wrap(t.put(
		ctx,
		t.filePath(articleID, file),
		file.MediaType,
		r,
	))
}

func (t *StorageArticleFile) PutThumbnail(
	ctx context.Context,
	articleID entity.ArticleID,
	file *entity.ArticleFileThumbnail,
	r io.Reader,
) error {
	return terrors.Wrap(t.put(
		ctx,
		t.filePathThumbnail(articleID, file),
		file.MediaType,
		r,
	))
}

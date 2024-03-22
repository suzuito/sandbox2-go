package infra

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *StorageArticle) GetArticle(ctx context.Context, articleID entity.ArticleID, w io.Writer) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(t.filePathMarkdown(articleID)).NewReader(ctx)
	if err != nil {
		return terrors.Wrap(err)
	}
	defer reader.Close()
	if _, err := io.Copy(w, reader); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

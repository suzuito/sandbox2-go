package bgcp

import (
	"context"
	"fmt"
	"io"
	"path"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog/entity"
)

func filePath(rootPath string, articleID entity.ArticleID, articleVersion int32) string {
	return path.Join(rootPath, string(articleID), fmt.Sprintf("%d.html", articleVersion))
}

type RepositoryArticleHTML struct {
	Cli      *storage.Client
	Bucket   string
	RootPath string
}

func (t *RepositoryArticleHTML) SetArticle(
	ctx context.Context,
	article *entity.Article,
	html string,
) error {
	writer := t.Cli.Bucket(t.Bucket).Object(filePath(t.RootPath, article.ID, article.Version)).NewWriter(ctx)
	_, err := writer.Write([]byte(html))
	if err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}

func (t *RepositoryArticleHTML) GetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	articleVersion int32,
	html io.Writer,
) error {
	reader, err := t.Cli.Bucket(t.Bucket).Object(filePath(t.RootPath, articleID, articleVersion)).NewReader(ctx)
	if err != nil {
		return err
	}

	_, err = io.Copy(html, reader)
	if err != nil {
		return err
	}
	return nil
}

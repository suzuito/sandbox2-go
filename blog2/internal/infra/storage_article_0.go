package infra

import (
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticle struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageArticle) filePathMarkdown(articleID entity.ArticleID) string {
	return fmt.Sprintf("%s.md", articleID)
}

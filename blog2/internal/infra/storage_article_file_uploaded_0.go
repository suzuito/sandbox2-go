package infra

import (
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticleFileUploaded struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageArticleFileUploaded) filePath(articleID entity.ArticleID, fileID entity.ArticleFileUploadedID) string {
	return fmt.Sprintf("%s/%s", articleID, fileID)
}

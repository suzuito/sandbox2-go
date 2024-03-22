package infra

import (
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type StorageArticleFileDirectlyUploaded struct {
	Cli    *storage.Client
	Bucket string
}

func (t *StorageArticleFileDirectlyUploaded) filePath(articleID entity.ArticleID, fileID entity.ArticleFileDirectlyUploadedID) string {
	return fmt.Sprintf("%s/%s", articleID, fileID)
}

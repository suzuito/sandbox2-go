package local

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
	"github.com/suzuito/sandbox2-go/internal/blog/usecase"
)

type RepositoryArticle struct {
}

func (t *RepositoryArticle) SearchArticles(
	ctx context.Context,
	query usecase.SearchArticlesQuery,
	articles *[]entity.Article,
) error {
	return &usecase.RepositoryError{
		EntityURL: "gcs://articles",
		Message:   "Not impl",
		Code:      usecase.RepositoryErrorCodeNotImpl,
	}
}

func (t *RepositoryArticle) SetArticle(
	ctx context.Context,
	article *entity.Article,
) error {
	b, err := json.Marshal(article)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%s.json", article.ID), b, 0600)
}

package usecase

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog/entity"
)

type RepositoryArticleHTML interface {
	SetArticle(
		ctx context.Context,
		article *entity.Article,
		html string,
	) error
	GetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		articleVersion int32,
		html io.Writer,
	) error
}

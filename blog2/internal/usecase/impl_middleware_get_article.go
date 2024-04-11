package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOMiddlewareGetArticle struct {
	Article *entity.Article
}

func (t *Impl) MiddlewareGetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*DTOMiddlewareGetArticle, error) {
	article, err := t.S.GetArticle(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOMiddlewareGetArticle{
		Article: article,
	}, nil
}

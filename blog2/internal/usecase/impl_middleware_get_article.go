package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/blog2/internal/usecase/pkg/serviceerror"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOMiddlewareGetArticle struct {
	Article *entity.Article
}

func (t *Impl) MiddlewareGetArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	publishedOnly bool,
) (*DTOMiddlewareGetArticle, error) {
	article, err := t.S.GetArticle(ctx, articleID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	if publishedOnly && !article.Published {
		return nil, terrors.Wrap(&serviceerror.NotFoundEntityError{
			EntityType: "articles",
			EntityID:   string(articleID),
		})
	}
	return &DTOMiddlewareGetArticle{
		Article: article,
	}, nil
}

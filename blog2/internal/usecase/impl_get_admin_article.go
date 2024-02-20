package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticle struct {
	Article *entity.Article
}

func (t *Impl) GetAdminArticle(
	ctx context.Context,
	articleID entity.ArticleID,
) (*DTOGetAdminArticle, error) {
	return nil, terrors.Wrapf("not impl")
}

package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticleTags struct {
	Tags []*entity.Tag
}

func (t *Impl) GetAdminArticleTags(
	ctx context.Context,
	article *entity.Article,
) (*DTOGetAdminArticleTags, error) {
	notAttachedArticleTags, err := t.S.GetNotAttachedArticleTags(ctx, article)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticleTags{
		Tags: notAttachedArticleTags,
	}, nil
}

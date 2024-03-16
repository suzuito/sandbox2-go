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
	// タグはあまり増えないという想定なので、全件取得で問題ない
	// TODO タグが増えてくると、パフォーマンスに影響が出てくる。その時はページング機構を実装すべき。
	tags, err := t.RepositoryArticle.GetAllTags(ctx)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	tagsNotAttached := []*entity.Tag{}
	for _, tag := range tags {
		if article.HasTag(tag.ID) {
			continue
		}
		tagsNotAttached = append(tagsNotAttached, tag)
	}
	return &DTOGetAdminArticleTags{
		Tags: tagsNotAttached,
	}, nil
}

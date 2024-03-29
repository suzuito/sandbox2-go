package infra

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) GetAllTags(ctx context.Context) ([]*entity.Tag, error) {
	// タグはあまり増えないという想定なので、全件取得で問題ない
	// TODO タグが増えてくると、パフォーマンスに影響が出てくる。その時はページング機構を実装すべき。
	rows, err := queryContext(
		ctx, t.Pool,
		"SELECT `id`, `name` FROM `tags`",
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rows.Close()
	tags := []*entity.Tag{}
	for rows.Next() {
		tag := entity.Tag{}
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, terrors.Wrap(err)
		}
		tags = append(tags, &tag)
	}
	return tags, nil
}

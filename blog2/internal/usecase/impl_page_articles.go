package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOPageArticles struct {
	Articles []*entity.Article
	PrevPage *int
	NextPage *int
}

func (t *Impl) PageArticles(
	ctx context.Context,
	tagID entity.TagID,
	page int,
	size int,
) (*DTOPageArticles, error) {
	offset := page * size
	publishedValueTrue := true
	var ptrTagID *entity.TagID
	if tagID != "" {
		ptrTagID = &tagID
	}
	q := entity.ArticleSearchQuery{
		ListQuery: entity.ListQuery{
			Offset: &offset,
			Limit:  &size,
		},
		Published: &publishedValueTrue,
		TagID:     ptrTagID,
	}
	articles, _, _, err := t.S.SearchArticles(ctx, &q)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	var ptrPrevPage *int
	var ptrNextPage *int
	if page > 0 {
		prevPage := page - 1
		ptrPrevPage = &prevPage
	}
	if len(articles) >= size {
		nextPage := page + 1
		ptrNextPage = &nextPage
	}
	return &DTOPageArticles{
		Articles: articles,
		PrevPage: ptrPrevPage,
		NextPage: ptrNextPage,
	}, nil
}

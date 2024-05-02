package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticles struct {
	Articles []*entity.Article
	NextPage *int
	PrevPage *int
}

func (t *Impl) GetAdminArticles(
	ctx context.Context,
	tagID entity.TagID,
	page int,
	size int,
	published *bool,
) (*DTOGetAdminArticles, error) {
	offset := page * size
	var ptrTagID *entity.TagID
	if tagID != "" {
		ptrTagID = &tagID
	}
	q := entity.ArticleSearchQuery{
		ListQuery: entity.ListQuery{
			Offset: &offset,
			Limit:  &size,
		},
		Published: published,
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
	return &DTOGetAdminArticles{
		Articles: articles,
		PrevPage: ptrPrevPage,
		NextPage: ptrNextPage,
	}, nil
}

type DTOPostAdminArticles struct {
	Article *entity.Article
}

func (t *Impl) PostAdminArticles(
	ctx context.Context,
) (*DTOPostAdminArticles, error) {
	article, err := t.S.CreateArticle(ctx)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOPostAdminArticles{
		Article: article,
	}, nil
}

package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

type DTOGetAdminArticles struct {
	Articles   []*entity.Article
	NextOffset *int
	PrevOffset *int
}

func (t *Impl) GetAdminArticles(
	ctx context.Context,
	query *entity.ArticleSearchQuery,
) (*DTOGetAdminArticles, error) {
	articles, prevOffset, nextOffset, err := t.S.SearchArticles(ctx, query)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &DTOGetAdminArticles{
		Articles:   articles,
		PrevOffset: prevOffset,
		NextOffset: nextOffset,
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

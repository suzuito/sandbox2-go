package infra

import (
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) GetArticles(ctx context.Context, ids ...entity.ArticleID) ([]*entity.Article, error) {
	return getArticles(ctx, t.Pool, ids...)
}

func getArticles(ctx context.Context, txOrDB TxOrDB, ids ...entity.ArticleID) ([]*entity.Article, error) {
	if len(ids) <= 0 {
		return []*entity.Article{}, nil
	}
	idsAsAny := toAnySlice(ids)

	// N+1問題に気を付けること
	// Batch Loadingによって回避

	// Get tags
	tags := map[entity.ArticleID][]entity.Tag{}
	rowsTags, err := queryContext(
		ctx,
		txOrDB,
		fmt.Sprintf(
			"SELECT `mapping_articles_tags`.`article_id` AS `article_id`, `tags`.`id` AS `tag_id`, `tags`.`name` AS `tag_name` "+
				"FROM `mapping_articles_tags` LEFT JOIN `tags` ON `mapping_articles_tags`.`tag_id` = `tags`.`id` WHERE %s",
			sqlIn(`article_id`, ids),
		),
		idsAsAny...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rowsTags.Close()
	for rowsTags.Next() {
		articleID := entity.ArticleID("")
		tag := entity.Tag{}
		if err := rowsTags.Scan(&articleID, &tag.ID, &tag.Name); err != nil {
			return nil, terrors.Wrap(err)
		}
		tags[articleID] = append(tags[articleID], tag)
	}

	// Get articles
	rowsArticle, err := queryContext(
		ctx,
		txOrDB,
		fmt.Sprintf(
			"SELECT `id`, `title`, `published`, `published_at` FROM `articles` WHERE %s",
			sqlIn(`id`, ids),
		),
		idsAsAny...,
	)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	defer rowsArticle.Close()
	articles := []*entity.Article{}
	for rowsArticle.Next() {
		article := entity.Article{}
		if err := rowsArticle.Scan(&article.ID, &article.Title, &article.Published, &article.PublishedAt); err != nil {
			return nil, terrors.Wrap(err)
		}
		tagsPerArticle, exists := tags[article.ID]
		if !exists {
			tagsPerArticle = []entity.Tag{}
		}
		article.Tags = tagsPerArticle
		articles = append(articles, &article)
	}

	return articles, nil
}

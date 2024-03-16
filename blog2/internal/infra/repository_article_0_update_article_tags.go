package infra

import (
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) UpdateArticleTags(
	ctx context.Context,
	articleID entity.ArticleID,
	add []entity.TagID,
	delete []entity.TagID,
) (*entity.Article, error) {
	if len(add) <= 0 && len(delete) <= 0 {
		return nil, terrors.Wrapf("must set at least one entity.TagID")
	}
	var article *entity.Article
	err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		if len(delete) > 0 {
			for i := range delete {
				query := "DELETE FROM `mapping_articles_tags` WHERE `article_id` = ? AND `tag_id` = ?"
				args := []any{articleID, delete[i]}
				_, err := execContext(ctx, tx, query, args...)
				if err != nil {
					return terrors.Wrap(err)
				}
			}
		}
		if len(add) > 0 {
			query := "INSERT IGNORE `mapping_articles_tags`(`article_id`, `tag_id`) VALUES "
			args := []any{}
			for i := range add {
				query += "(?,?)"
				args = append(args, articleID, add[i])
				if i < len(add)-1 {
					query += ","
				}
			}
			_, err := execContext(ctx, tx, query, args...)
			if err != nil {
				return terrors.Wrap(err)
			}
		}
		articles, err := getArticles(ctx, tx, articleID)
		if err != nil {
			return terrors.Wrap(err)
		}
		if len(articles) <= 0 {
			return terrors.Wrap(fmt.Errorf("inconsistency DB state. try to update tags of unexist article %s", articleID))
		}
		article = articles[0]
		return nil
	})
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return article, nil
}

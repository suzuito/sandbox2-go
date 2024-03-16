package infra

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *RepositoryArticle) UpdateArticle(
	ctx context.Context,
	articleID entity.ArticleID,
	title *string,
	published *bool,
	publishedAt *time.Time,
) (*entity.Article, error) {
	if title == nil && published == nil && publishedAt == nil {
		return nil, terrors.Wrap(fmt.Errorf("all fields are null"))
	}
	var article *entity.Article
	err := withTransaction(ctx, t.Pool, func(tx TxOrDB) error {
		args := []any{}
		q := "UPDATE `articles` SET "
		statsSet := []string{}
		if title != nil {
			statsSet = append(statsSet, "`articles`.`title` = ?")
			args = append(args, *title)
		}
		if published != nil {
			statsSet = append(statsSet, "`articles`.`published` = ?")
			args = append(args, *published)
		}
		if publishedAt != nil {
			statsSet = append(statsSet, "`articles`.`published_at` = ?")
			args = append(args, *publishedAt)
		}
		q += strings.Join(statsSet, ",")
		q += " WHERE `articles`.`id` = ?"
		args = append(args, articleID)
		_, err := execContext(ctx, tx, q, args...)
		if err != nil {
			return terrors.Wrap(err)
		}

		articles, err := getArticles(ctx, tx, articleID)
		if err != nil {
			return terrors.Wrap(err)
		}
		article = articles[0]
		return nil
	})
	if err != nil {
		return nil, terrors.Wrap(err)
	}

	return article, nil
}

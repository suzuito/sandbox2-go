package infra

import (
	"context"
	"fmt"
	"strings"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	"github.com/suzuito/sandbox2-go/common/csql"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func updateSearchIndex(ctx context.Context, txOrDB csql.TxOrDB, articleID entity.ArticleID) error {
	articles, err := getArticles(ctx, txOrDB, articleID)
	if err != nil {
		return terrors.Wrap(err)
	}
	if len(articles) <= 0 {
		return terrors.Wrap(fmt.Errorf("document of articleID '%s' is not found", articleID))
	}
	article := articles[0]
	tagIDs := []string{}
	for _, tagID := range article.GetTagIDs() {
		tagIDs = append(tagIDs, string(tagID))
	}
	_, err = txOrDB.ExecContext(
		ctx,
		"INSERT INTO `articles_search_index`(`article_id`, `tags`, `published`, `created_at`, `published_at`) VALUES(?, ?, ?, ?, ?)"+
			" ON DUPLICATE KEY UPDATE `tags` = ?, `published` = ?, `created_at` = ?, `published_at` = ?",
		article.ID,
		strings.Join(tagIDs, " "),
		article.Published,
		article.CreatedAt,
		article.PublishedAt,
		strings.Join(tagIDs, " "),
		article.Published,
		article.CreatedAt,
		article.PublishedAt,
	)
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

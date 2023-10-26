package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type SearchArticlesQuery struct {
	SortField SearchArticlesQuerySortField
	Offset    int
	Limit     int
	Tags      []string
	SortOrder SortOrder
}

type SearchArticlesQuerySortField string

var (
	SearchArticlesQuerySortFieldDate    SearchArticlesQuerySortField = "date"
	SearchArticlesQuerySortFieldVersion SearchArticlesQuerySortField = "version"
)

type RepositoryArticle interface {
	GetLatestArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		article *entity.Article,
	) error
	GetArticleByPrimaryKey(
		ctx context.Context,
		primaryKey entity.ArticlePrimaryKey,
		article *entity.Article,
	) error
	GetArticlesByPrimaryKey(
		ctx context.Context,
		primaryKeys []entity.ArticlePrimaryKey,
		sortField SearchArticlesQuerySortField,
		sortOrder SortOrder,
		articles *[]entity.Article,
	) error
	GetArticlesByID(
		ctx context.Context,
		articleID entity.ArticleID,
		articles *[]entity.Article,
	) error
	GetArticlesByArticleSourceID(
		ctx context.Context,
		articleSourceID entity.ArticleSourceID,
		articles *[]entity.Article,
	) error
	SetArticle(
		ctx context.Context,
		article *entity.Article,
	) error
	SetArticleSearchIndex(
		ctx context.Context,
		article *entity.Article,
	) error
	SearchArticles(
		ctx context.Context,
		query SearchArticlesQuery,
		articlePrimaryKeys *[]entity.ArticlePrimaryKey,
		hasNext *bool,
	) error
}

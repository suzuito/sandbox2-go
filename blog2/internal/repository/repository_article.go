package repository

import (
	"context"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryArticle interface {
	GetArticles(ctx context.Context, ids ...entity.ArticleID) ([]*entity.Article, error)
	CreateArticle(ctx context.Context, articleID entity.ArticleID) (*entity.Article, error)
	UpdateArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
		published *bool,
		publishedAt *time.Time,
	) (*entity.Article, error)

	GetAllTags(ctx context.Context) ([]*entity.Tag, error)
	UpdateArticleTags(
		ctx context.Context,
		articleID entity.ArticleID,
		add []entity.TagID,
		delete []entity.TagID,
	) (*entity.Article, error)
}

package repository

import (
	"context"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type RepositoryArticle interface {
	GetArticles(ctx context.Context, ids ...entity.ArticleID) ([]*entity.Article, error)
	SearchArticles(ctx context.Context, q *entity.ArticleSearchQuery) ([]*entity.Article, *int, *int, error)
	CreateArticle(ctx context.Context, articleID entity.ArticleID, createdAt time.Time) (*entity.Article, error)
	UpdateArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
		published *bool,
		publishedAt *time.Time,
	) (*entity.Article, error)

	GetAllTags(ctx context.Context) ([]*entity.Tag, error)
	CreateTag(ctx context.Context, name string) (*entity.Tag, error)
	UpdateArticleTags(
		ctx context.Context,
		articleID entity.ArticleID,
		add []entity.TagID,
		delete []entity.TagID,
	) (*entity.Article, error)

	PutFile(
		ctx context.Context,
		file *entity.File,
	) error
	GetFile(
		ctx context.Context,
		fileID entity.FileID,
	) (*entity.File, error)
	DeleteFile(
		ctx context.Context,
		fileID entity.FileID,
	) error
	SearchFiles(
		ctx context.Context,
		queryString string,
		offset int,
		limit int,
	) ([]*entity.FileAndThumbnail, error)

	PutFileThumbnail(
		ctx context.Context,
		file *entity.FileThumbnail,
	) error
	GetFileThumbnail(
		ctx context.Context,
		fileID entity.FileID,
	) (*entity.FileThumbnail, error)
	DeleteFileThumbnail(
		ctx context.Context,
		fileID entity.FileID,
	) error
}

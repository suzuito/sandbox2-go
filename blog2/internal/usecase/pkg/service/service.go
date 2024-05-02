package service

import (
	"context"
	"io"
	"time"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type Service interface {
	GetNotAttachedArticleTags(ctx context.Context, article *entity.Article) ([]*entity.Tag, error)
	GetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*entity.Article, error)
	GetArticleBody(
		ctx context.Context,
		articleID entity.ArticleID,
	) (string, string, error)
	PutArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
		published *bool,
		publishedAt *time.Time,
	) (*entity.Article, error)
	PutArticleMarkdown(
		ctx context.Context,
		articleID entity.ArticleID,
		markdownBody string,
	) (string, error)
	UpdateArticleTags(
		ctx context.Context,
		articleID entity.ArticleID,
		add []entity.TagID,
		delete []entity.TagID,
	) (*entity.Article, error)
	SearchArticles(ctx context.Context, q *entity.ArticleSearchQuery) ([]*entity.Article, *int, *int, error)
	CreateArticle(
		ctx context.Context,
	) (*entity.Article, error)
	CreateFileUploaded(
		ctx context.Context,
		fileName string,
		fileType entity.FileType,
		input io.Reader,
	) (*entity.FileUploaded, error)

	DetectFileType(
		ctx context.Context,
		data []byte,
	) (entity.FileType, string)
	PutFile(
		ctx context.Context,
		file *entity.File,
		data []byte,
	) error
	ExistFile(
		ctx context.Context,
		fileID entity.FileID,
	) error

	CreateThumbnail(
		ctx context.Context,
		fileID entity.FileID,
	) (*entity.FileThumbnail, error)

	SearchFiles(
		ctx context.Context,
		queryString string,
		offset int,
		limit int,
	) ([]*entity.FileAndThumbnail, error)

	CreateTestData(ctx context.Context) error
}

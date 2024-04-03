package usecase

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
	internal_usecase "github.com/suzuito/sandbox2-go/blog2/internal/usecase"
)

type Usecase interface {
	GetAdminArticles(
		ctx context.Context,
		query *entity.ArticleSearchQuery,
	) (*internal_usecase.DTOGetAdminArticles, error)
	PostAdminArticles(
		ctx context.Context,
	) (*internal_usecase.DTOPostAdminArticles, error)
	GetAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*internal_usecase.DTOGetAdminArticle, error)
	PutAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
	) (*internal_usecase.DTOPutAdminArticle, error)
	PutAdminArticleMarkdown(
		ctx context.Context,
		articleID entity.ArticleID,
		markdownBodyReader io.Reader,
	) error
	PostAdminArticlePublish(
		ctx context.Context,
		articleID entity.ArticleID,
	) error
	DeleteAdminArticlePublish(
		ctx context.Context,
		articleID entity.ArticleID,
	) error
	PostAdminArticleEditTags(
		ctx context.Context,
		articleID entity.ArticleID,
		add []entity.TagID,
		delete []entity.TagID,
	) error
	GetAdminArticleTags(
		ctx context.Context,
		article *entity.Article,
	) (*internal_usecase.DTOGetAdminArticleTags, error)
	PostAdminArticleImages(
		ctx context.Context,
		article *entity.Article,
		input io.Reader,
	) (*internal_usecase.DTOPostAdminArticleImages, error)
	MiddlewareGetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*internal_usecase.DTOMiddlewareGetArticle, error)

	StartImageProcessFromGCF(ctx context.Context, data []byte) error

	// Not production codes
	CreateTestData001(
		ctx context.Context,
	) error
}

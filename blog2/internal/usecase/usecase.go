package usecase

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/blog2/internal/entity"
)

type Usecase interface {
	GetAdminArticles(
		ctx context.Context,
		query *entity.ArticleSearchQuery,
	) (*DTOGetAdminArticles, error)
	PostAdminArticles(
		ctx context.Context,
	) (*DTOPostAdminArticles, error)
	GetAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*DTOGetAdminArticle, error)
	PutAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
	) (*DTOPutAdminArticle, error)
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
	) (*DTOGetAdminArticleTags, error)
	PostAdminArticleImages(
		ctx context.Context,
		article *entity.Article,
		input io.Reader,
	) (*DTOPostAdminArticleImages, error)
	MiddlewareGetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*DTOMiddlewareGetArticle, error)
}

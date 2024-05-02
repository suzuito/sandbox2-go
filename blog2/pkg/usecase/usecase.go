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
		article *entity.Article,
	) (*internal_usecase.DTOGetAdminArticle, error)
	PutAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
		published *bool,
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
	// PostAdminFiles(
	// 	ctx context.Context,
	// 	fileName string,
	// 	fileType entity.FileType,
	// 	input io.Reader,
	// ) (*internal_usecase.DTOPostAdminFiles, error)
	MiddlewareGetArticle(
		ctx context.Context,
		articleID entity.ArticleID,
	) (*internal_usecase.DTOMiddlewareGetArticle, error)

	APIPutAdminArticle(
		ctx context.Context,
		articleID entity.ArticleID,
		title *string,
		published *bool,
	) (*internal_usecase.DTOAPIPutAdminArticle, error)
	APIPostAdminArticleEditTags(
		ctx context.Context,
		article *entity.Article,
		add []entity.TagID,
		delete []entity.TagID,
	) (*internal_usecase.DTOAPIPostAdminArticleEditTags, error)
	APIPutAdminArticleMarkdown(
		ctx context.Context,
		articleID entity.ArticleID,
		markdown io.Reader,
	) (*internal_usecase.DTOAPIPutAdminArticleMarkdown, error)
	APIPostAdminFiles(
		ctx context.Context,
		fileName string,
		file io.Reader,
	) (*internal_usecase.DTOAPIPostAdminFiles, error)
	APIGetAdminFiles(
		ctx context.Context,
		queryString string,
		page int,
		limit int,
	) (*internal_usecase.DTOAPIGetAdminFiles, error)

	// Not production codes
	CreateTestData001(
		ctx context.Context,
	) error
}

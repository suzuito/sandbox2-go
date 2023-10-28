package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/blog/usecase/markdown2html"
)

type UsecaseImpl struct {
	RepositoryArticle       RepositoryArticle
	RepositoryArticleSource RepositoryArticleSource
	RepositoryArticleHTML   RepositoryArticleHTML
	Markdown2HTML           markdown2html.Markdown2HTML
}

type Usecase interface {
	GenerateArticleHTMLFromMarkdown(
		ctx context.Context,
		articleSource *entity.ArticleSource,
		md []byte,
	) (*entity.Article, string, error)
	GenerateArticleHTML(
		ctx context.Context,
		articleSourceID entity.ArticleSourceID,
		articleSourceVersion string,
	) (*entity.Article, string, error)
	GenerateSitemap(
		ctx context.Context,
		siteOrigin string,
		urls *XMLURLSet,
	) error
	SearchArticles(
		ctx context.Context,
		query SearchArticlesQuery,
		articles *[]entity.Article,
		hasNext *bool,
	) error
	UploadArticle(
		ctx context.Context,
		articleSourceID entity.ArticleSourceID,
		articleSourceVersion string,
	) (*entity.Article, error)
	UploadAllArticles(
		ctx context.Context,
		ref string,
	) error
}

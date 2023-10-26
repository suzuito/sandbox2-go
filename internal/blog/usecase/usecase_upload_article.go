package usecase

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

func (t *UsecaseImpl) uploadArticle(
	ctx context.Context,
	articleSource *entity.ArticleSource,
	md []byte,
) (*entity.Article, error) {
	// t.L.Infof("Upload article source '%s'\n", articleSource.ID)
	article, html, err := t.GenerateArticleHTMLFromMarkdown(
		ctx,
		articleSource,
		md,
	)
	if err != nil {
		return nil, err
	}
	if err := t.RepositoryArticleHTML.SetArticle(ctx, article, html); err != nil {
		return nil, err
	}
	if err := t.RepositoryArticle.SetArticle(ctx, article); err != nil {
		return nil, err
	}
	if err := t.RepositoryArticle.SetArticleSearchIndex(ctx, article); err != nil {
		return nil, err
	}
	return article, nil
}

func (t *UsecaseImpl) UploadArticle(
	ctx context.Context,
	articleSourceID entity.ArticleSourceID,
	articleSourceVersion string,
) (*entity.Article, error) {
	articleSource, src, err := t.RepositoryArticleSource.GetArticleSource(ctx, articleSourceID, articleSourceVersion)
	if err != nil {
		return nil, err
	}
	return t.uploadArticle(
		ctx,
		articleSource,
		src,
	)
}

// TODO このメソッドを消す。
func (t *UsecaseImpl) UploadAllArticles(
	ctx context.Context,
	ref string,
) error {
	return t.RepositoryArticleSource.GetArticleSources(
		ctx,
		ref,
		func(
			articleSource *entity.ArticleSource,
			articleSourceBody []byte,
		) error {
			_, err := t.uploadArticle(ctx, articleSource, articleSourceBody)
			return err
		},
	)
}

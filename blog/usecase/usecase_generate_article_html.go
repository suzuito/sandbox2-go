package usecase

import (
	"context"
	"fmt"

	"github.com/suzuito/sandbox2-go/blog/entity"
	"github.com/suzuito/sandbox2-go/common/terrors"
)

func (t *UsecaseImpl) GenerateArticleHTMLFromMarkdown(
	ctx context.Context,
	articleSource *entity.ArticleSource,
	md []byte,
) (*entity.Article, string, error) {
	article := entity.Article{
		ArticleSource: *articleSource,
	}
	articleHTML := ""
	if err := t.Markdown2HTML.Generate(ctx, string(md), &articleHTML, &article); err != nil {
		return nil, "", err
	}
	if err := article.Validate(); err != nil {
		return nil, "", terrors.Wrap(fmt.Errorf("Invalid article : %w", err))
	}
	return &article, articleHTML, nil
}

func (t *UsecaseImpl) GenerateArticleHTML(
	ctx context.Context,
	articleSourceID entity.ArticleSourceID,
	articleSourceVersion string,
) (*entity.Article, string, error) {
	articleSource, md, err := t.RepositoryArticleSource.GetArticleSource(
		ctx,
		articleSourceID,
		articleSourceVersion,
	)
	if err != nil {
		return nil, "", err
	}
	return t.GenerateArticleHTMLFromMarkdown(ctx, articleSource, md)
}

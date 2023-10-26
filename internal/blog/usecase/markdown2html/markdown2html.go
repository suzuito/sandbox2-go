package markdown2html

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/blog/entity"
)

type Markdown2HTML interface {
	Generate(
		ctx context.Context,
		src string,
		dst *string,
		article *entity.Article,
	) error
}

type Markdown2HTMLImpl struct{}

func (t *Markdown2HTMLImpl) Generate(
	ctx context.Context,
	src string,
	dst *string,
	article *entity.Article,
) error {
	if err := convertMarkdownToHTML(ctx, src, dst, article); err != nil {
		return err
	}
	if err := extractArticleTitleFromHTML(ctx, *dst, article); err != nil {
		return err
	}
	return nil
}

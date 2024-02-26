package markdown2html

import (
	"context"

	"github.com/suzuito/sandbox2-go/blog/entity"
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
) error {
	// Pre processor
	// Markdown to HTML
	if err := convertMarkdownToHTML(ctx, src, dst); err != nil {
		return err
	}
	// Post processor
	return nil
}

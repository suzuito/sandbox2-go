package markdown2html

import (
	"context"
)

type Markdown2HTML interface {
	Generate(
		ctx context.Context,
		src string,
		dst *string,
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

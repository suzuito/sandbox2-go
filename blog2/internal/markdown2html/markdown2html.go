package markdown2html

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
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
	html := ""
	if err := convertMarkdownToHTML(ctx, src, &html); err != nil {
		return terrors.Wrap(err)
	}
	// Post processor
	if err := convertHTML(ctx, html, dst); err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

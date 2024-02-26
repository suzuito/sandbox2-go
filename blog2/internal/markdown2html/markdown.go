package markdown2html

import (
	"bytes"
	"context"
	"fmt"

	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/toc"
)

type astTransformerAddLinkBlank struct {
}

func (a *astTransformerAddLinkBlank) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() != ast.KindLink && n.Kind() != ast.KindAutoLink {
			return ast.WalkContinue, nil
		}
		n.SetAttributeString("target", []byte("blank"))
		return ast.WalkContinue, nil
	})
}

type astTransformerHeading struct {
}

func (a *astTransformerHeading) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() != ast.KindHeading {
			return ast.WalkContinue, nil
		}
		n.SetAttributeString("class", []byte("md-heading"))
		return ast.WalkContinue, nil
	})
}

func convertMarkdownToHTML(
	ctx context.Context,
	src string,
	dst *string,
) error {
	fmt.Println(src)
	buffer := bytes.NewBufferString("")
	parserContext := parser.NewContext()
	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,         // Parse yaml meta
			extension.Linkify, // Convert a link in markdown to a <a> tag in html
			highlighting.NewHighlighting(
				highlighting.WithStyle("github"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
					chromahtml.WithCustomCSS(map[chroma.TokenType]string{
						chroma.PreWrapper: "overflow: scroll;",
					}),
				),
			), // Syntax highlight
			emoji.Emoji,
			&toc.Extender{},
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Can write raw HTML in Markdown
		),
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&astTransformerAddLinkBlank{}, 1), // Add target="_blank" to <a> tag in html
				util.Prioritized(&astTransformerHeading{}, 1),
			),
			parser.WithAutoHeadingID(), // For TOC
		),
	)
	if err := md.Convert(
		[]byte(src),
		buffer,
		parser.WithContext(parserContext),
	); err != nil {
		return err
	}
	*dst = buffer.String()
	return nil
}

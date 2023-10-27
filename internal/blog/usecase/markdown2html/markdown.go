package markdown2html

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/suzuito/sandbox2-go/internal/blog/entity"
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
	article *entity.Article,
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

	metaMap := meta.Get(parserContext)
	fmt.Printf("%+v\n", metaMap)
	article.Description = extractMetaValue(metaMap, "description", "")
	dateAtString := extractMetaValue(metaMap, "date", "")
	date, err := time.Parse("2006-01-02", dateAtString)
	if err != nil {
		article.Date = time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	} else {
		article.Date = date
	}
	tagsString := extractMetaValue(metaMap, "tags", []interface{}{})
	article.Version = int32(extractMetaValue(metaMap, "version", int(0)))
	article.Tags = []entity.Tag{}
	for _, tagString := range tagsString {
		article.Tags = append(article.Tags, entity.Tag{ID: entity.TagID(tagString.(string))})
	}
	article.ID = entity.ArticleID(extractMetaValue(metaMap, "id", ""))
	fmt.Printf("%+v\n", article)
	return nil
}

func extractMetaValue[T any](metaMap map[string]interface{}, key string, dflt T) T {
	v := metaMap[key]
	if v == nil {
		return dflt
	}
	vv, ok := v.(T)
	if !ok {
		return dflt
	}
	return vv
}

package markdown2html

import (
	"bytes"
	"context"
	"strconv"

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
)

// Goldmarkの内部処理の概要
// https://github.com/yuin/goldmark?tab=readme-ov-file#goldmark-internalfor-extension-developers

// GoldmarkのTransformer
// Transformerは、ASTを走査し、Nodeを変換するためのインターフェース

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

type astTransformerAddMarkdownLines struct {
}

func (a *astTransformerAddMarkdownLines) Transform(
	node *ast.Document,
	reader text.Reader,
	pc parser.Context,
) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Type() != ast.TypeBlock {
			return ast.WalkContinue, nil
		}
		if entering {
			if n.Kind() != ast.KindListItem && n.HasChildren() {
				for child := n.FirstChild(); ; child = child.NextSibling() {
					if child == n.LastChild() {
						break
					}
				}
			}
			lines := getLinesInMarkdown(n, reader)
			if lines >= 0 {
				n.SetAttributeString("data-source-line", []byte(strconv.Itoa(lines)))
			}
			return ast.WalkContinue, nil
		}
		if n.Kind() == ast.KindListItem && n.HasChildren() {
			// 後で解説が必要
			// ulタグにdata-source-line属性がつかないので
			v, ok := n.FirstChild().Attribute([]byte("data-source-line"))
			if ok {
				n.SetAttribute([]byte("data-source-line"), v)
			}
		}
		return ast.WalkContinue, nil
	})
}

func getLinesInMarkdown(node ast.Node, reader text.Reader) int {
	if node.Lines().Len() <= 0 {
		return -1
	}
	segment := node.Lines().At(node.Lines().Len() - 1)
	segment.Start = 0
	lines := 1
	for _, b := range string(reader.Value(segment)) {
		if b == '\n' {
			// fmt.Println(b)
			lines++
		}
	}
	// fmt.Println("====")
	// fmt.Println(string(reader.Value(segment)))
	// fmt.Println("=========>", lines, strings.Count(string(reader.Value(segment)), "\n"))
	// fmt.Println(strings.ReplaceAll(string(reader.Value(segment)), "\n", "<???>"))
	// fmt.Println("++++++++++")
	return lines
}

type astTransformerAddImageInfo struct {
}

func (a *astTransformerAddImageInfo) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	ast.Walk(node, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() != ast.KindImage {
			return ast.WalkContinue, nil
		}
		if !entering {
			return ast.WalkContinue, nil
		}
		nodeImage, ok := n.(*ast.Image)
		if !ok {
			return ast.WalkContinue, nil
		}
		nodeImage.SetAttributeString("style", []byte("max-width: 100%;"))
		nodeImage.SetAttributeString("class", []byte("article-image")) // Postprocessorにて、リンク付き画像にするためにclassを追加する
		return ast.WalkContinue, nil
	})
}

func convertMarkdownToHTML(
	ctx context.Context,
	src string,
	dst *string,
) error {
	// fmt.Println(src)
	buffer := bytes.NewBufferString("")
	parserContext := parser.NewContext()
	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,         // Parse yaml meta
			extension.Linkify, // Convert a link in markdown to a <a> tag in html
			highlighting.NewHighlighting(
				highlighting.WithStyle("vim"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
					chromahtml.WithLinkableLineNumbers(true, "L"),
					chromahtml.WithCustomCSS(map[chroma.TokenType]string{
						chroma.PreWrapper: "overflow: scroll; padding: 10px;",
					}),
				),
			), // Syntax highlight
			emoji.Emoji,
			// &toc.Extender{}, // Generate TOC
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(), // Can write raw HTML in Markdown
		),
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&astTransformerAddLinkBlank{}, 1), // Add target="_blank" to <a> tag in html
				util.Prioritized(&astTransformerHeading{}, 1),
				util.Prioritized(&astTransformerAddMarkdownLines{}, 1),
				util.Prioritized(&astTransformerAddImageInfo{}, 1),
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

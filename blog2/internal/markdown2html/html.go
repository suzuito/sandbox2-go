package markdown2html

import (
	"bytes"
	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func convertHTML(ctx context.Context, src string, dst *string) error {
	r := bytes.NewBufferString(src)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return terrors.Wrap(err)
	}
	// 画像をリンク付き画像へ変換する
	doc.Find(".article-image").Each(func(i int, s *goquery.Selection) {
		if s.Nodes[0].DataAtom != atom.Img {
			return
		}
		imgnode := s.Nodes[0]
		srcURL := ""
		for _, attr := range imgnode.Attr {
			if attr.Key == "src" {
				srcURL = attr.Val
			}
		}
		if srcURL == "" {
			return
		}
		anode := html.Node{
			Type:      html.ElementNode,
			DataAtom:  atom.A,
			Data:      "a",
			Namespace: "",
			Attr: []html.Attribute{
				{Key: "href", Val: srcURL},
				{Key: "target", Val: "_blank"},
			},
		}
		parent := imgnode.Parent
		parent.RemoveChild(imgnode)
		anode.AppendChild(imgnode)
		parent.AppendChild(&anode)
	})
	*dst, err = doc.Html()
	if err != nil {
		return terrors.Wrap(err)
	}
	return nil
}

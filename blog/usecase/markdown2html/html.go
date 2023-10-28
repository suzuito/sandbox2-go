package markdown2html

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/blog/entity"
)

func extractArticleTitleFromHTML(
	ctx context.Context,
	src string,
	article *entity.Article,
) error {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(src))
	if err != nil {
		return err
	}
	article.Title = d.Find("h1.md-heading").First().Text()
	return nil
}

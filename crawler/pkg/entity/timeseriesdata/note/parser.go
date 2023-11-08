package note

import (
	"context"
	"io"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/constant"
)

func subString(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

// Parser for article on note (note.com)
// ex) https://note.com/knowledgework/n/n46b7881a16a6
type Parser struct{}

func (t *Parser) Parse(
	ctx context.Context,
	r io.Reader,
) (*TimeSeriesDataNoteArticle, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	article := TimeSeriesDataNoteArticle{}
	// Title
	selTitle := doc.Find("title").First()
	if selTitle == nil {
		return nil, terrors.Wrapf("Cannot find title tag")
	}
	if len(selTitle.Text()) <= 0 {
		return nil, terrors.Wrapf("Cannot find title tag")
	}
	article.Title = selTitle.Text()
	// URL
	selCanonical := doc.Find("link[rel=canonical]")
	if selCanonical.Length() <= 0 {
		return nil, terrors.Wrapf("Cannot find link[rel=canonical] tag")
	}
	hrefURLString, exists := selCanonical.First().Attr("href")
	if !exists {
		return nil, terrors.Wrapf("Cannot find href attr of link[rel=canonical] tag")
	}
	article.URL = hrefURLString
	// ImageURL
	article.ImageURL = doc.Find("meta[property='og:image']").First().AttrOr("content", "")
	// ArticleContent
	selArticleContent := doc.Find(".p-article__content")
	if selArticleContent.Length() <= 0 {
		return nil, terrors.Wrapf("Cannot find html tag of .p-article__content")
	}
	article.ArticleContent = selArticleContent.First().Text()
	// Description
	selDescription := doc.Find("meta[name=description]").First()
	if selDescription == nil {
		article.Description = subString(article.ArticleContent, 100)
	} else {
		article.Description = selDescription.AttrOr("content", subString(article.ArticleContent, 100))
	}
	// PublishedAt
	selArticlePublishedAt := doc.Find(".o-noteContentHeader__info time")
	if selArticlePublishedAt.Length() <= 0 {
		return nil, terrors.Wrapf("Cannot find html tag of '.o-noteContentHeader__info time'")
	}
	publishedAtString, exists := selArticlePublishedAt.First().Attr("datetime")
	if !exists {
		return nil, terrors.Wrapf("Cannot find datetime attr in html tag of '.o-noteContentHeader__info time'")
	}
	publishedAt, err := time.Parse(time.RFC3339, publishedAtString)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	article.PublishedAt = publishedAt.In(constant.JST)
	// Tags
	selTags := doc.Find(".m-tagList__item")
	selTags.Each(func(i int, s *goquery.Selection) {
		txt := s.Text()
		replaced := regexp.MustCompile(`\s+`).ReplaceAll([]byte(txt), []byte{})
		replaced = regexp.MustCompile(`^#`).ReplaceAll(replaced, []byte{})
		article.Tags = append(article.Tags, TimeSeriesDataNoteArticleTag{
			Name: string(replaced),
		})
	})
	return &article, nil
}

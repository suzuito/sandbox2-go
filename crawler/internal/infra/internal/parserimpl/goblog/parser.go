package goblog

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Parser struct {
	BaseURLGoBlog string
}

func (t *Parser) ID() crawler.ParserID {
	return "goblog"
}

func (t *Parser) Do(ctx context.Context, r io.Reader, _ crawler.CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	returned := []timeseriesdata.TimeSeriesData{}
	doc.Find(".blogtitle").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		if title == "" {
			return
		}
		href, exists := s.Find("a").Attr("href")
		if !exists {
			return
		}
		dateString := s.Find(".date").Text()
		if dateString == "" {
			return
		}
		blogURL := fmt.Sprintf("%s%s", t.BaseURLGoBlog, href)
		publishedAt, err := time.Parse("2 January 2006", dateString)
		if err != nil {
			clog.L.Errorf(ctx, "%+v", terrors.Wrap(err))
			return
		}
		data := timeseriesdata.TimeSeriesDataBlogFeed{
			ID:          timeseriesdata.TimeSeriesDataID(fmt.Sprintf("goblog-%s", publishedAt.Format("2006-01-02"))),
			PublishedAt: publishedAt,
			Title:       title,
			URL:         blogURL,
			Author: &timeseriesdata.TimeSeriesDataBlogFeedAuthor{
				Name:     "goblog",
				URL:      "https://go.dev",
				ImageURL: "https://go.dev/images/favicon-gopher.png",
			},
		}
		returned = append(returned, &data)
	})
	return returned, nil
}

func New(def *crawler.ParserDefinition, _ *factorysetting.CrawlerFactorySetting) (crawler.Parser, error) {
	parser := Parser{
		BaseURLGoBlog: "https://go.dev",
	}
	if def.ID != parser.ID() {
		return nil, factory.ErrNoMatchedParserID
	}
	return &parser, nil
}

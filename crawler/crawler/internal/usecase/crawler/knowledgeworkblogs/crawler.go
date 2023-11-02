package knowledgeworkblogs

import (
	"context"
	"net/http"

	"github.com/mmcdole/gofeed"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

const CrawlerID crawler.CrawlerID = "knowledgeworkblog"

type Crawler struct {
}

func NewCrawler() crawler.Crawler {
	return &Crawler{}
}

func (t *Crawler) ID() crawler.CrawlerID {
	return CrawlerID
}

func (t *Crawler) Name() string {
	return string(CrawlerID)
}

func (t *Crawler) NewFetcher(ctx context.Context) (crawler.Fetcher, error) {
	return &Fetcher{
		cliHTTP: http.DefaultClient,
	}, nil
}

func (t *Crawler) NewParser(ctx context.Context) (crawler.Parser, error) {
	return &Parser{
		fp: gofeed.NewParser(),
	}, nil
}

func (t *Crawler) NewPublisher(ctx context.Context) (crawler.Publisher, error) {
	return &Publisher{}, nil
}

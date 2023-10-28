package goblog

import (
	"context"
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
)

const CrawlerID crawler.CrawlerID = "goblog"
const baseURLGoBlog = "https://go.dev"

type Crawler struct {
	repository repository.Repository
}

func NewCrawler(repository repository.Repository) crawler.Crawler {
	return &Crawler{
		repository: repository,
	}
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
	return &Parser{}, nil
}

func (t *Crawler) NewPublisher(ctx context.Context) (crawler.Publisher, error) {
	return &Publisher{
		repository: t.repository,
	}, nil
}

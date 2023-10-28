package crawlerfactory

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
)

type CrawlerFactory interface {
	GetCrawler(ctx context.Context, crawlerID crawler.CrawlerID) (crawler.Crawler, error)
	GetCrawlers(ctx context.Context, crawlerIDs ...crawler.CrawlerID) []crawler.Crawler
}

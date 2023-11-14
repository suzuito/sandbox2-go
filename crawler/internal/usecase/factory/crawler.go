package factory

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type CrawlerFactory interface {
	Get(ctx context.Context, def *crawler.CrawlerDefinition) (*crawler.Crawler, error)
}

package crawler

import (
	"context"
)

type CrawlerID string

type Crawler interface {
	ID() CrawlerID
	Name() string

	NewFetcher(ctx context.Context) (Fetcher, error)
	NewParser(ctx context.Context) (Parser, error)
	NewPublisher(ctx context.Context) (Publisher, error)
}

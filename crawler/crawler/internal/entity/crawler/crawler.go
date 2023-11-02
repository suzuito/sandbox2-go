package crawler

import (
	"context"
	"io"
)

type CrawlerID string

type Crawler interface {
	ID() CrawlerID
	Name() string

	ParseInput(ctx context.Context, input io.Reader) (Input, error)
	NewFetcher(ctx context.Context) (Fetcher, error)
	NewParser(ctx context.Context) (Parser, error)
	NewPublisher(ctx context.Context) (Publisher, error)
}

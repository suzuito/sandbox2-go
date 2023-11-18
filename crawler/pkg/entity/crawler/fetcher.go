package crawler

import (
	"context"
	"io"
)

type FetcherID string

type Fetcher interface {
	ID() FetcherID
	Do(ctx context.Context, w io.Writer, input CrawlerInputData) error
}

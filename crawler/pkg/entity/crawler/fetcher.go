package crawler

import (
	"context"
	"io"
)

type FetcherID string

type Fetcher interface {
	Do(ctx context.Context, w io.Writer, input CrawlerInputData) error
}

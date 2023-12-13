package crawler

import (
	"context"
	"io"
	"log/slog"
)

type FetcherID string

type Fetcher interface {
	ID() FetcherID
	Do(ctx context.Context, logger *slog.Logger, w io.Writer, input CrawlerInputData) error
}

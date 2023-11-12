package crawler

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Crawler2 struct {
	Fetcher   Fetcher
	Parser    Parser
	Publisher Publisher
}

type Fetcher interface {
	Do(ctx context.Context, w io.Writer, input CrawlerInputData) error
}

type Parser interface {
	Do(ctx context.Context, r io.Writer, input CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error)
}

type Publisher interface {
	Do(ctx context.Context, input CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error
}

package crawler

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type ParserID string

type Parser interface {
	Do(ctx context.Context, r io.Reader, input CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error)
}

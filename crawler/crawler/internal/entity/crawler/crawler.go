package crawler

import (
	"context"
	"io"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type CrawlerID string

type Crawler interface {
	ID() CrawlerID
	Name() string

	Fetch(ctx context.Context, w io.Writer, msg CrawlerInputData) error
	Parse(ctx context.Context, r io.Reader, msg CrawlerInputData) ([]timeseriesdata.TimeSeriesData, error)
	Publish(ctx context.Context, msg CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error
}

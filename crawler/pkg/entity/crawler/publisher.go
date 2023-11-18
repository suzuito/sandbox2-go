package crawler

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type PublisherID string

type Publisher interface {
	ID() PublisherID
	Do(ctx context.Context, input CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error
}

package publisher

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Publisher interface {
	Do(ctx context.Context, input crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error
}

package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type TimeSeriesDataRepository interface {
	SetTimeSeriesData(
		ctx context.Context,
		crawlerID crawler.CrawlerID,
		timeSeriesData ...timeseriesdata.TimeSeriesData,
	) error
}

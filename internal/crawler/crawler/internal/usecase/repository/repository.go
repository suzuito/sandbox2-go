package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/internal/crawler/pkg/entity/timeseriesdata"
)

type Repository interface {
	SetTimeSeriesData(
		ctx context.Context,
		crawlerID crawler.CrawlerID,
		timeSeriesData ...timeseriesdata.TimeSeriesData,
	) error
}

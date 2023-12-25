package publisherimpl

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type TimeSeriesDataRepositoryPublisher struct {
	Repository           repository.TimeSeriesDataRepository
	TimeSeriesDataBaseID timeseriesdata.TimeSeriesDataBaseID
}

func (t *TimeSeriesDataRepositoryPublisher) ID() crawler.PublisherID {
	return "timeseriesdatarepository"
}

func (t *TimeSeriesDataRepositoryPublisher) Do(ctx context.Context, input crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.Repository.SetTimeSeriesData(ctx, t.TimeSeriesDataBaseID, data...))
}

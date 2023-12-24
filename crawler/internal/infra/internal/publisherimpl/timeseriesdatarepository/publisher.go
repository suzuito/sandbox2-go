package timeseriesdatarepository

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type Publisher struct {
	Repository           repository.TimeSeriesDataRepository
	TimeSeriesDataBaseID timeseriesdata.TimeSeriesDataBaseID
}

func (t *Publisher) ID() crawler.PublisherID {
	return "timeseriesdatarepository"
}

func (t *Publisher) Do(ctx context.Context, input crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.Repository.SetTimeSeriesData(ctx, t.TimeSeriesDataBaseID, data...))
}

func New(def *crawler.PublisherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Publisher, error) {
	publisher := Publisher{
		Repository: setting.PublisherFactorySetting.TimeSeriesDataRepository,
	}
	if def.ID != publisher.ID() {
		return nil, factory.ErrNoMatchedPublisherID
	}
	var err error
	publisher.TimeSeriesDataBaseID, err = argument.GetFromArgumentDefinition[timeseriesdata.TimeSeriesDataBaseID](def.Argument, "TimeSeriesDataBaseID")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &publisher, nil
}

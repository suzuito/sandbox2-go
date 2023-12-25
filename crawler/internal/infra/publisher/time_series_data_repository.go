package publisher

import (
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/publisher/internal/publisherimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

func NewTimeSeriesDataRepositoryPublisher(def *crawler.PublisherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Publisher, error) {
	publisher := publisherimpl.TimeSeriesDataRepositoryPublisher{
		Repository: setting.PublisherFactorySetting.TimeSeriesDataRepository,
	}
	if def.ID != publisher.ID() {
		return nil, factoryerror.ErrNoMatchedPublisherID
	}
	var err error
	publisher.TimeSeriesDataBaseID, err = argument.GetFromArgumentDefinition[timeseriesdata.TimeSeriesDataBaseID](def.Argument, "TimeSeriesDataBaseID")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	return &publisher, nil
}

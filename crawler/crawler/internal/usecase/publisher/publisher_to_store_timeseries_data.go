package publisher

import (
	"context"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/entity/crawler"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/timeseriesdata"
)

type RepositoryToStoreTimeseriesData struct {
	repository repository.Repository
	crawlerID  crawler.CrawlerID
}

func (t *RepositoryToStoreTimeseriesData) Do(ctx context.Context, _ crawler.CrawlerInputData, data ...timeseriesdata.TimeSeriesData) error {
	return terrors.Wrap(t.repository.SetTimeSeriesData(ctx, t.crawlerID, data...))
}

func NewRepositoryToStoreTimeseriesData(
	repository repository.Repository,
	crawlerID crawler.CrawlerID,
) *RepositoryToStoreTimeseriesData {
	return &RepositoryToStoreTimeseriesData{
		repository: repository,
		crawlerID:  crawlerID,
	}
}

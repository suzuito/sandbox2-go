package inject

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawlerdefinitions"
)

func NewUsecaseLocal(ctx context.Context) (usecase.Usecase, error) {
	var env Environment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	projectID := "dummy-prj"
	fcli, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	pcli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	httpClient := http.DefaultClient
	u := usecase.UsecaseImpl{
		L:                        clog.L,
		TriggerCrawlerQueue:      infra.NewTriggerCrawlerQueue(pcli, "gcf-CrawlerCrawl"),
		CrawlerRepository:        infra.NewCrawlerRepository(crawlerdefinitions.AvailableCrawlers),
		CrawlerFactory:           infra.NewCrawlerFactory(httpClient),
		TimeSeriesDataRepository: infra.NewTimeSeriesDataRepository(fcli, "Crawler"),
	}
	return &u, nil
}

package internal

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/infra/gcp"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/infra/web"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/crawler/crawler/pkg/inject"
)

func NewUsecaseLocal(ctx context.Context) (*usecase.UsecaseImpl, error) {
	var env inject.Environment
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
	repository := gcp.NewRepository(fcli, "Crawler")
	queue := gcp.NewQueue(pcli, "gcf-CrawlerCrawl")
	u := usecase.UsecaseImpl{
		Repository: repository,
		Queue:      queue,
		CrawlerFactory: crawlerfactory.NewDefaultCrawlerFactoryImpl(
			repository,
			queue,
			web.NewFetcherHTTP(http.DefaultClient),
		),
		L: clog.L,
	}
	return &u, nil
}

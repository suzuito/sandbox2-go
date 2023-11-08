package inject

import (
	"context"
	"net/http"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/infra/gcp"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/infra/web"
	internal_usecase "github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase"
	"github.com/suzuito/sandbox2-go/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/crawler/crawler/pkg/usecase"
)

func NewUsecaseGCP(ctx context.Context) (usecase.Usecase, error) {
	projectID, err := metadata.ProjectID()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	var env Environment
	err = envconfig.Process("", &env)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	fcli, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	pcli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	repository := gcp.NewRepository(fcli, "Crawler")
	queue := gcp.NewQueue(pcli, "gcf-CrawlerCrawl")
	u := internal_usecase.UsecaseImpl{
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

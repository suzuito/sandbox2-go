package inject

import (
	"context"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/infra/gcp"
	internal_usecase "github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/pkg/usecase"
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
	repository := gcp.NewRepository(fcli, "Crawler")
	u := internal_usecase.UsecaseImpl{
		Repository:     gcp.NewRepository(fcli, "Crawler"),
		Queue:          gcp.NewQueue(pcli, "gcf-CrawlerCrawl"),
		CrawlerFactory: crawlerfactory.NewDefaultCrawlerFactoryImpl(repository),
		L:              clog.L,
	}
	return &u, nil
}

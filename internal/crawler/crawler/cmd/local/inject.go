package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/infra/gcp"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/internal/usecase/crawlerfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/crawler/pkg/inject"
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
	u := usecase.UsecaseImpl{
		Repository:     gcp.NewRepository(fcli, "Crawler"),
		Queue:          gcp.NewQueue(pcli, "gcf-CrawlerCrawl"),
		CrawlerFactory: crawlerfactory.NewDefaultCrawlerFactoryImpl(repository),
		L:              clog.L,
	}
	return &u, nil
}

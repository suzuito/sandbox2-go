package inject

import (
	"context"
	"log/slog"
	"net/http"
	"os"

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
	slogHandler := clog.CustomHandler{
		Handler: slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
			},
		),
	}
	logger := slog.New(&slogHandler)
	timeSeriesDataRepository := infra.NewTimeSeriesDataRepository(fcli, "Crawler")
	httpClient := http.DefaultClient
	u := usecase.UsecaseImpl{
		L:                        logger,
		TriggerCrawlerQueue:      infra.NewTriggerCrawlerQueue(pcli, "gcf-CrawlerCrawl"),
		CrawlerRepository:        infra.NewCrawlerRepository(crawlerdefinitions.AvailableCrawlers),
		CrawlerFactory:           infra.NewCrawlerFactory(httpClient, timeSeriesDataRepository),
		TimeSeriesDataRepository: timeSeriesDataRepository,
	}
	return &u, nil
}

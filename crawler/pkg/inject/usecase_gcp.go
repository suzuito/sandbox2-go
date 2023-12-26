package inject

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	infra_factory "github.com/suzuito/sandbox2-go/crawler/internal/infra/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"
	infra_queue "github.com/suzuito/sandbox2-go/crawler/internal/infra/queue"
	infra_repository "github.com/suzuito/sandbox2-go/crawler/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/crawler/internal/inject"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase"
	pkg_usecase "github.com/suzuito/sandbox2-go/crawler/pkg/usecase"
)

func NewUsecaseGCP(ctx context.Context) (pkg_usecase.Usecase, error) {
	projectID, err := metadata.ProjectID()
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	var env inject.Environment
	if err := envconfig.Process("", &env); err != nil {
		return nil, terrors.Wrap(err)
	}
	fcli, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	firestoreBaseCollection := "Crawler"
	pcli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	discordGoSession, err := discordgo.New("Bot " + env.GoVillageDiscordBotToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	discordGoSession.LogLevel = discordgo.LogDebug
	discordGoSession.Debug = true
	slogHandler := clog.CustomHandler{
		Handler: slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelInfo,
				AddSource: true,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Key == slog.LevelKey {
						a.Key = "severity"
					}
					return a
				},
			},
		),
	}
	logger := slog.New(&slogHandler)
	timeSeriesDataRepository := infra_repository.NewTimeSeriesDataRepository(fcli, firestoreBaseCollection)
	triggerCrawlerQueue := infra_queue.NewTriggerCrawlerQueue(
		pcli,
		"gcf-CrawlerCrawl",
		"gcf-CrawlerDispatchCrawl",
	)
	httpClient := http.DefaultClient
	u := usecase.UsecaseImpl{
		L:                              logger,
		TriggerCrawlerQueue:            triggerCrawlerQueue,
		CrawlerRepository:              infra_repository.NewCrawlerRepository(inject.AvailableCrawlers, inject.CrawlerStarterSettings),
		CrawlerConfigurationRepository: infra_repository.NewCrawlerConfigurationRepository(),
		CrawlerFactory: infra_factory.NewCrawlerFactory(&factorysetting.CrawlerFactorySetting{
			FetcherFactorySetting: factorysetting.FetcherFactorySetting{
				HTTPClientWrapper: httpclientwrapper.NewHTTPClientWrapper(
					httpClient,
					infra_repository.NewHTTPClientCacheRepository(fcli, firestoreBaseCollection),
				),
			},
			ParserFactorySetting: factorysetting.ParserFactorySetting{},
			PublisherFactorySetting: factorysetting.PublisherFactorySetting{
				TriggerCrawlerQueue:      triggerCrawlerQueue,
				TimeSeriesDataRepository: timeSeriesDataRepository,
			},
		}),
		NotifierRepository: infra_repository.NewNotifierRepository(inject.NewAvailableNotifiers(&env)),
		NotifierFactory: infra_factory.NewNotifierFactory(&factorysetting.NotifierFactorySetting{
			DiscordClient: discordGoSession,
		}),
		TimeSeriesDataRepository: timeSeriesDataRepository,
	}
	return &u, nil
}

package inject

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/httpclientcache/gcpimpl"
	"github.com/suzuito/sandbox2-go/common/terrors"
	infra_factory "github.com/suzuito/sandbox2-go/crawler/internal/infra/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	infra_queue "github.com/suzuito/sandbox2-go/crawler/internal/infra/queue"
	infra_repository "github.com/suzuito/sandbox2-go/crawler/internal/infra/repository"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase"
	pkg_usecase "github.com/suzuito/sandbox2-go/crawler/pkg/usecase"
)

func NewUsecaseLocal(ctx context.Context) (pkg_usecase.Usecase, error) {
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
	firestoreBaseCollection := "Crawler"
	pcli, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	scli, err := storage.NewClient(ctx)
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
		Handler: slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level:     slog.LevelDebug,
				AddSource: true,
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
		CrawlerRepository:              infra_repository.NewCrawlerRepository(AvailableCrawlers, CrawlerStarterSettings),
		CrawlerConfigurationRepository: infra_repository.NewCrawlerConfigurationRepository(),
		CrawlerFactory: infra_factory.NewCrawlerFactory(&factorysetting.CrawlerFactorySetting{
			FetcherFactorySetting: factorysetting.FetcherFactorySetting{
				HTTPClient: httpClient,
				HTTPClientCacheClient: gcpimpl.New(
					scli,
					env.HTTPClientCacheBucket,
					env.HTTPClientCacheBasePath,
				),
			},
			ParserFactorySetting: factorysetting.ParserFactorySetting{},
			PublisherFactorySetting: factorysetting.PublisherFactorySetting{
				TriggerCrawlerQueue:      triggerCrawlerQueue,
				TimeSeriesDataRepository: timeSeriesDataRepository,
			},
		}),
		NotifierRepository: infra_repository.NewNotifierRepository(NewAvailableNotifiers(&env)),
		NotifierFactory: infra_factory.NewNotifierFactory(&factorysetting.NotifierFactorySetting{
			DiscordClient: discordGoSession,
		}),
		TimeSeriesDataRepository: timeSeriesDataRepository,
	}
	return &u, nil
}

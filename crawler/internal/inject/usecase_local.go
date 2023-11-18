package inject

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra"
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
	triggerCrawlerQueue := infra.NewTriggerCrawlerQueue(pcli, "gcf-CrawlerCrawl")
	httpClient := http.DefaultClient
	u := usecase.UsecaseImpl{
		L:                        logger,
		TriggerCrawlerQueue:      triggerCrawlerQueue,
		CrawlerRepository:        infra.NewCrawlerRepository(AvailableCrawlers),
		CrawlerFactory:           infra.NewCrawlerFactory(httpClient, timeSeriesDataRepository, triggerCrawlerQueue),
		NotifierRepository:       infra.NewNotifierRepository(NewAvailableNotifiers(&env)),
		NotifierFactory:          infra.NewNotifierFactory(discordGoSession),
		TimeSeriesDataRepository: timeSeriesDataRepository,
	}
	return &u, nil
}

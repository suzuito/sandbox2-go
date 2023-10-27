package inject

import (
	"context"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/infra/gcp"
	internal_usecase "github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifierfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/pkg/usecase"
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
	discordGoSession, err := discordgo.New("Bot " + env.GoVillageDiscordBotToken)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	fcli, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	u := internal_usecase.UsecaseImpl{
		L:          clog.L,
		Repository: gcp.NewRepository(fcli),
		NotifierFactory: notifierfactory.NewDefaultNotifierFactoryImpl(
			discordGoSession,
			env.GoVillageDiscordChannelIDNews,
			env.GoVillageDiscordChannelIDEvents,
		),
	}
	return &u, nil
}

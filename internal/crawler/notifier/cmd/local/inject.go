package main

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/common/terrors"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/infra/gcp"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifierfactory"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/pkg/inject"
)

func NewUsecaseLocal(ctx context.Context) (*usecase.UsecaseImpl, error) {
	projectID := "dummy-prj"
	var env inject.Environment
	if err := envconfig.Process("", &env); err != nil {
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
	u := usecase.UsecaseImpl{
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

package factory

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/notifierimpl/discordblogfeed"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/notifierimpl/discordconnpassevent"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/discord"
	usecase_factory "github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
)

func NewNotifierFactory(
	DiscordClient discord.DiscordGoSession,
) usecase_factory.NotifierFactory {
	return &factory.NotifierFactory{
		DiscordClient: DiscordClient,
		NewFuncs: []factory.NewFuncNotifier{
			discordconnpassevent.New,
			discordblogfeed.New,
		},
	}
}

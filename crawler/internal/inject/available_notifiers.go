package inject

import (
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifierdefinitions/discordblogfeed"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifierdefinitions/discordconnpassevent"
)

func NewAvailableNotifiers(env *Environment) []notifier.NotifierDefinition {
	return []notifier.NotifierDefinition{
		*discordblogfeed.New(env.GoVillageDiscordChannelIDNews),
		*discordconnpassevent.New(env.GoVillageDiscordChannelIDEvents),
	}
}

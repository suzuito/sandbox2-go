package discordconnpassevent

import (
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

func New(discordChannelID string) *notifier.NotifierDefinition {
	return &notifier.NotifierDefinition{
		ID: "discordconnpassevent",
		DocPathFirestoreMatchers: []string{
			`Crawler/TimeSeriesData/goconnpass/.+$`,
		},
		Argument: argument.ArgumentDefinition{
			"DiscordChannelID": discordChannelID,
		},
	}
}

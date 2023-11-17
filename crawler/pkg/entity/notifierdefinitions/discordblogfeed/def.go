package discordblogfeed

import (
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

func New(discordChannelID string) *notifier.NotifierDefinition {
	return &notifier.NotifierDefinition{
		ID: "discordblogfeed",
		DocPathFirestoreMatchers: []string{
			`Crawler/TimeSeriesData/goblog/.+$`,
			`Crawler/TimeSeriesData/golangweekly/.+$`,
			`Crawler/TimeSeriesData/knowledgeworkblog/.+$`,
		},
		Argument: argument.ArgumentDefinition{
			"DiscordChannelID": discordChannelID,
		},
	}
}

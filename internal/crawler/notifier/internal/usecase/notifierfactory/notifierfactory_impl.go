package notifierfactory

import (
	"context"
	"regexp"

	"github.com/suzuito/sandbox2-go/internal/common/cusecase/clog"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/entity/notifier"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifier/govillageblogfeed"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifier/govillageconnpassevent"
	"github.com/suzuito/sandbox2-go/internal/crawler/notifier/internal/usecase/notifier/govillagegolangweekly"
)

type NotifierFactoryImpl struct {
	entries []NotifierEntry
}

func (t *NotifierFactoryImpl) GetNotiferFromDocPathFirestore(
	ctx context.Context,
	fullPath string,
) ([]notifier.Notifier, error) {
	matchedNotifiers := []notifier.Notifier{}
	for _, entry := range t.entries {
		for _, matcher := range entry.DocPathFirestoreMatchers {
			matched, err := regexp.MatchString(matcher, fullPath)
			if err != nil {
				clog.L.Errorf(ctx, "%+v", err)
				continue
			}
			if !matched {
				continue
			}
			matchedNotifiers = append(matchedNotifiers, entry.Notifier)
		}
	}
	return matchedNotifiers, nil
}

func newNotifierFactoryImpl(entries []NotifierEntry) *NotifierFactoryImpl {
	factory := NotifierFactoryImpl{
		entries: entries,
	}
	return &factory
}

func NewDefaultNotifierFactoryImpl(
	discordGoSession discord.DiscordGoSession,
	goVillageDiscordChannelIDNews string,
	goVillageDiscordChannelIDEvents string,
) *NotifierFactoryImpl {
	return newNotifierFactoryImpl(
		[]NotifierEntry{
			{
				DocPathFirestoreMatchers: []string{
					`Crawler/TimeSeriesData/goblog/.+$`,
				},
				Notifier: govillageblogfeed.NewNotifier(
					discordGoSession,
					goVillageDiscordChannelIDNews,
				),
			},
			{
				DocPathFirestoreMatchers: []string{
					`Crawler/TimeSeriesData/goconnpass/.+$`,
				},
				Notifier: govillageconnpassevent.NewNotifier(
					discordGoSession,
					goVillageDiscordChannelIDEvents,
				),
			},
			{
				DocPathFirestoreMatchers: []string{
					`Crawler/TimeSeriesData/golangweekly/.+$`,
				},
				Notifier: govillagegolangweekly.NewNotifier(
					discordGoSession,
					goVillageDiscordChannelIDNews,
				),
			},
		},
	)
}

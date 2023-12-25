package notifier

import (
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/notifier/internal/notifierimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

func NewDiscordBlogFeedNotifier(def *notifier.NotifierDefinition, setting *factorysetting.NotifierFactorySetting) (notifier.Notifier, error) {
	n := notifierimpl.DiscordBlogFeedNotifier{
		DiscordClient: setting.DiscordClient,
	}
	if def.ID != n.ID() {
		return nil, factoryerror.ErrNoMatchedNotifierID
	}
	discordChannelID, err := argument.GetFromArgumentDefinition[string](def.Argument, "DiscordChannelID")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	n.DiscordChannelID = discordChannelID
	return &n, nil
}

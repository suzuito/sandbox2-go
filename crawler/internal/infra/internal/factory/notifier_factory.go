package factory

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/discord"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

type NewFuncNotifierArgument struct {
	DiscordClient discord.DiscordGoSession
}
type NewFuncNotifier func(def *notifier.NotifierDefinition, arg *NewFuncNotifierArgument) (notifier.Notifier, error)

type NotifierFactory struct {
	DiscordClient discord.DiscordGoSession
	NewFuncs      []NewFuncNotifier
}

func (t *NotifierFactory) Get(ctx context.Context, def *notifier.NotifierDefinition) (notifier.Notifier, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, &NewFuncNotifierArgument{
			DiscordClient: t.DiscordClient,
		})
		if err != nil {
			if errors.Is(err, ErrNoMatchedNotifierID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Notifier '%s' is not found in available list", def.ID)
}

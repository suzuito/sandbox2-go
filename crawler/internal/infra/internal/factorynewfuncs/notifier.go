package factorynewfuncs

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/notifierimpl/discordblogfeed"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/notifierimpl/discordconnpassevent"
)

var NewFuncsNotifier = []factory.NewFuncNotifier{
	discordconnpassevent.New,
	discordblogfeed.New,
}

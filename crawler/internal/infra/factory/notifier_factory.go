package factory

import (
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/internal/factoryimpl"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/notifier"
	usecase_factory "github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
)

func NewNotifierFactory(
	setting *factorysetting.NotifierFactorySetting,
) usecase_factory.NotifierFactory {
	return &factoryimpl.NotifierFactory{
		NotifierFactorySetting: setting,
		NewFuncs: []factoryimpl.NewFuncNotifier{
			notifier.NewDiscordBlogFeedNotifier,
			notifier.NewDiscordConnpassEventNotifier,
		},
	}
}

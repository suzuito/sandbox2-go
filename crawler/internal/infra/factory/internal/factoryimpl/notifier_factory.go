package factoryimpl

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

type NewFuncNotifier func(def *notifier.NotifierDefinition, setting *factorysetting.NotifierFactorySetting) (notifier.Notifier, error)

type NotifierFactory struct {
	NotifierFactorySetting *factorysetting.NotifierFactorySetting
	NewFuncs               []NewFuncNotifier
}

func (t *NotifierFactory) Get(ctx context.Context, def *notifier.NotifierDefinition) (notifier.Notifier, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, t.NotifierFactorySetting)
		if err != nil {
			if errors.Is(err, factoryerror.ErrNoMatchedNotifierID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Notifier '%s' is not found in available list", def.ID)
}

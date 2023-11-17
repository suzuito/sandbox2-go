package factory

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

type NotifierFactory interface {
	Get(ctx context.Context, def *notifier.NotifierDefinition) (notifier.Notifier, error)
}

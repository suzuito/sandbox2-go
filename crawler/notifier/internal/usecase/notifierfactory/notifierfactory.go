package notifierfactory

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/notifier/internal/entity/notifier"
)

type NotifierEntry struct {
	DocPathFirestoreMatchers []string
	Notifier                 notifier.Notifier
}

type NotifierFactory interface {
	GetNotiferFromDocPathFirestore(
		ctx context.Context,
		fullPath string,
	) ([]notifier.Notifier, error)
}

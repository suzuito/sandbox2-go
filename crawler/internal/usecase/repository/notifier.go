package repository

import (
	"context"

	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/notifier"
)

type NotifierRepository interface {
	GetNotiferDefinitionsFromDocPathFirestore(
		ctx context.Context,
		fullPath string,
	) ([]notifier.NotifierDefinition, error)
}

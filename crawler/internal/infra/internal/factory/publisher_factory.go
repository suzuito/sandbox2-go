package factory

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/usecase/repository"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncPublisherArgument struct {
	TimeSeriesDataRepository repository.TimeSeriesDataRepository
}
type NewFuncPublisher func(def *crawler.PublisherDefinition, arg *NewFuncPublisherArgument) (crawler.Publisher, error)

type PublisherFactory struct {
	NewFuncs                 []NewFuncPublisher
	TimeSeriesDataRepository repository.TimeSeriesDataRepository
}

func (t *PublisherFactory) Get(ctx context.Context, def *crawler.PublisherDefinition) (crawler.Publisher, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, &NewFuncPublisherArgument{
			TimeSeriesDataRepository: t.TimeSeriesDataRepository,
		})
		if err != nil {
			if errors.Is(err, ErrNoMatchedFetcherID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Publisher '%s' is not found in available list", def.ID)
}

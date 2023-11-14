package factory

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncPublisherArgument struct {
}
type NewFuncPublisher func(def *crawler.PublisherDefinition, arg *NewFuncPublisherArgument) (crawler.Publisher, error)

type PublisherFactory struct {
	NewFuncs []NewFuncPublisher
}

func (t *PublisherFactory) Get(ctx context.Context, def *crawler.PublisherDefinition) (crawler.Publisher, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, &NewFuncPublisherArgument{})
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

package factory

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/pkg/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncPublisher func(def *crawler.PublisherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Publisher, error)

type PublisherFactory struct {
	CrawlerFactorySetting *factorysetting.CrawlerFactorySetting
	NewFuncs              []NewFuncPublisher
}

func (t *PublisherFactory) Get(ctx context.Context, def *crawler.PublisherDefinition) (crawler.Publisher, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, t.CrawlerFactorySetting)
		if err != nil {
			if errors.Is(err, ErrNoMatchedPublisherID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Publisher '%s' is not found in available list", def.ID)
}

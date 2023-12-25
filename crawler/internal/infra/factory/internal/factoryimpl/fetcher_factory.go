package factoryimpl

import (
	"context"
	"errors"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncFetcher func(def *crawler.FetcherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error)

var NewFuncsFetcher = []NewFuncFetcher{
	fetcher.NewFetcherHTTP,
	fetcher.NewFetcherHTTPStatic,
	fetcher.NewFetcherHTTPConnpass,
}

type FetcherFactory struct {
	CrawlerFactorySetting *factorysetting.CrawlerFactorySetting
	NewFuncs              []NewFuncFetcher
}

func (t *FetcherFactory) Get(ctx context.Context, def *crawler.FetcherDefinition) (crawler.Fetcher, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, t.CrawlerFactorySetting)
		if err != nil {
			if errors.Is(err, factoryerror.ErrNoMatchedFetcherID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Fetcher '%s' is not found in available list", def.ID)
}

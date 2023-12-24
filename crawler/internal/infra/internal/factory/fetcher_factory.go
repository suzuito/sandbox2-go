package factory

import (
	"context"
	"errors"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/httprequestcache"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type NewFuncFetcherArgument struct {
	HTTPClient       *http.Client
	HTTPRequestCache httprequestcache.HTTPRequestCache
}
type NewFuncFetcher func(def *crawler.FetcherDefinition, arg *NewFuncFetcherArgument) (crawler.Fetcher, error)

type FetcherFactory struct {
	HTTPClient *http.Client
	NewFuncs   []NewFuncFetcher
}

func (t *FetcherFactory) Get(ctx context.Context, def *crawler.FetcherDefinition) (crawler.Fetcher, error) {
	for _, newFunc := range t.NewFuncs {
		f, err := newFunc(def, &NewFuncFetcherArgument{
			HTTPClient: t.HTTPClient,
		})
		if err != nil {
			if errors.Is(err, ErrNoMatchedFetcherID) {
				continue
			}
			return nil, terrors.Wrap(err)
		}
		return f, nil
	}
	return nil, terrors.Wrapf("Fetcher '%s' is not found in available list", def.ID)
}

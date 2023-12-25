package fetcher

import (
	"net/url"
	"time"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/internal/fetcherimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewFetcherHTTPConnpass(def *crawler.FetcherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
	f := fetcherimpl.FetcherHTTPConnpass{}
	if f.ID() != def.ID {
		return nil, factoryerror.ErrNoMatchedFetcherID
	}
	days, err := argument.GetFromArgumentDefinition[int](def.Argument, "Days")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	query, err := argument.GetFromArgumentDefinition[url.Values](def.Argument, "Query")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	f.TimeNowFunc = time.Now
	f.Days = days
	f.Query = query
	f.Cli = setting.FetcherFactorySetting.HTTPClient
	return &f, nil
}

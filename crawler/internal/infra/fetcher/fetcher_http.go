package fetcher

import (
	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/internal/fetcherimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewFetcherHTTP(def *crawler.FetcherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
	f := fetcherimpl.FetcherHTTP{}
	if f.ID() != def.ID {
		return nil, factoryerror.ErrNoMatchedFetcherID
	}
	statusCodesSuccess, err := argument.GetFromArgumentDefinition[[]int](def.Argument, "StatusCodesSuccess")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	cli := httpclientwrapper.NewHTTPClientWrapperFromArgument(def, setting)
	f.Cli = cli
	f.StatusCodesSuccess = statusCodesSuccess
	return &f, nil
}

package fetcher

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/common/terrors"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factoryerror"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/internal/fetcherimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

func NewFetcherHTTPStatic(def *crawler.FetcherDefinition, setting *factorysetting.CrawlerFactorySetting) (crawler.Fetcher, error) {
	f := fetcherimpl.FetcherHTTPStatic{}
	if f.ID() != def.ID {
		return nil, factoryerror.ErrNoMatchedFetcherID
	}
	urlString, err := argument.GetFromArgumentDefinition[string](def.Argument, "URL")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	method, err := argument.GetFromArgumentDefinition[string](def.Argument, "Method")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	statusCodesSuccess, err := argument.GetFromArgumentDefinition[[]int](def.Argument, "StatusCodesSuccess")
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	req, err := http.NewRequest(method, urlString, nil)
	if err != nil {
		return nil, terrors.Wrap(err)
	}
	cli := httpclientwrapper.NewHTTPClientWrapperFromArgument(def, setting)
	f.Cli = cli
	f.Req = req
	f.StatusCodesSuccess = statusCodesSuccess
	return &f, nil
}

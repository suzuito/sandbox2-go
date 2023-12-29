package httpclientwrapper

import (
	"context"
	"io"
	"log/slog"
	"net/http"

	"github.com/suzuito/sandbox2-go/common/httpclientcache"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/factory/factorysetting"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/fetcher/httpclientwrapper/internal/httpclientwrapperimpl"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/argument"
	"github.com/suzuito/sandbox2-go/crawler/pkg/entity/crawler"
)

type HTTPClientWrapper interface {
	Do(
		ctx context.Context,
		logger *slog.Logger,
		req *http.Request,
		w io.Writer,
		statusCodesSuccess []int,
	) error
}

func NewHTTPClientWrapperFromArgument(
	def *crawler.FetcherDefinition,
	setting *factorysetting.CrawlerFactorySetting,
) HTTPClientWrapper {
	useCache := argument.DefaultGetFromArgumentDefinition[bool](def.Argument, "UseCache", false)
	httpClientCacheOption := argument.DefaultGetFromArgumentDefinition[*httpclientcache.ClientOption](def.Argument, "HTTPClientCacheOption", nil)
	return &httpclientwrapperimpl.HTTPClientWrapperImpl{
		Cli:         setting.FetcherFactorySetting.HTTPClient,
		UseCache:    useCache,
		Cache:       setting.FetcherFactorySetting.HTTPClientCacheClient,
		CacheOption: httpClientCacheOption,
	}
}

package infra

import (
	"net/http"

	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factory"
	"github.com/suzuito/sandbox2-go/crawler/internal/infra/internal/factorynewfuncs"
	usecase_factory "github.com/suzuito/sandbox2-go/crawler/internal/usecase/factory"
)

func NewCrawlerFactory(
	httpClient *http.Client,
) usecase_factory.CrawlerFactory {
	returned := factory.CrawlerFactory{
		FetcherFactory: factory.FetcherFactory{
			HTTPClient: httpClient,
			NewFuncs:   factorynewfuncs.NewFuncsFetcher,
		},
		ParserFactory: factory.ParserFactory{
			NewFuncs: factorynewfuncs.NewFuncsParser,
		},
		PublisherFactory: factory.PublisherFactory{
			NewFuncs: factorynewfuncs.NewFuncsPublisher,
		},
	}
	return &returned
}
